package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"transactor-server/pkg/account"
	"transactor-server/pkg/api"
	"transactor-server/pkg/infra/config"
	"transactor-server/pkg/infra/log"
	"transactor-server/pkg/metric"
	"transactor-server/pkg/operationtype"
	"transactor-server/pkg/tracer"
	"transactor-server/pkg/transaction"

	appconfig "transactor-server/pkg/config"
	"transactor-server/pkg/db"
	_ "transactor-server/pkg/docs"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/oklog/run"
	"go.uber.org/zap"
)

var flagConfig = flag.String("config", "./config/base.yml", "path to the config file")

func main() {
	flag.Parse()

	var cfg appconfig.Config

	// load config from file and override from ENV with APP_ prefix eg. APP_SERVER_PORT
	err := config.New(&cfg, config.FromFile(*flagConfig), config.FromENV("APP"))
	if err != nil {
		log.L.Fatal("", zap.Error(err))
	}

	var otelConn *grpc.ClientConn

	if cfg.Server.EnableTelemetry {
		otelConn, err = grpc.NewClient(cfg.Server.OTELEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.L.Fatal("", zap.Error(err))
		}
		defer otelConn.Close()
	}

	otelRes, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", appconfig.AppName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.L.Sugar().Fatalf("Could not set resources: %v", err)
	}

	if cfg.Server.EnableTelemetry {
		cleanup := tracer.Init(cfg.Server.OTELEndpoint, otelRes)
		defer cleanup(context.Background())
	}

	if cfg.Server.EnableTelemetry {
		cleanup := metric.Init(cfg.Server.OTELEndpoint, otelRes)
		defer cleanup(context.Background())
	}

	// setup a new logger for testing or prod
	// this is the global logger
	// for each component a custom logger with addititional attribute should be used
	logger, otlpLogSync := log.New(cfg.Server.Debug, otelConn, otelRes)
	defer logger.Sync()
	if otlpLogSync != nil {
		defer otlpLogSync.Close()
	}

	logger.Info("starting api server")

	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	entClient, err := db.OpenEntClient(ctx, cfg.DB)
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	defer entClient.Close()

	operationTypeDAO := operationtype.NewDAO(entClient)

	transactionDAO := transaction.NewDAO(entClient)
	transactionService := transaction.NewService(
		operationTypeDAO,
		transactionDAO,
		logger.With(zap.String("layer", "application"), zap.String("service", "transaction")),
	)
	transactionAPI := transaction.NewAPI(transactionService)

	accountDAO := account.NewDAO(entClient)
	accountService := account.NewService(
		accountDAO,
		logger.With(zap.String("layer", "application"), zap.String("service", "account")),
	)
	accountService = account.NewTracedService(accountService, logger.With(zap.String("layer", "application"), zap.String("service", "account")))
	accountService = account.NewMeteredService(accountService)
	accountAPI := account.NewAPI(accountService)

	app := api.NewRouter(cfg.Server.APIKey, transactionAPI, accountAPI, logger)

	var g run.Group
	{
		g.Add(func() error {
			logger.Info("server", zap.String("msg", "serving http"), zap.String("addr", addr))
			return app.Listen(addr)
		}, func(error) {
			logger.Info("server", zap.String("msg", "stopping http server"))
			if err := app.Shutdown(); err != nil {
				logger.Fatal("", zap.Error(err))
			}

			logger.Info("db", zap.String("msg", "stopping db"))
			if err := entClient.Close(); err != nil {
				logger.Fatal("", zap.Error(err))
			}
		})
	}
	{
		// set-up our signal handler
		var (
			cancelInterrupt = make(chan struct{})
			c               = make(chan os.Signal, 2)
		)
		defer close(c)

		g.Add(func() error {
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}

	logger.Error("exit", zap.Error(g.Run()))
}
