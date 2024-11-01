package api

import (
	"fmt"
	"net/http"
	"transactor-server/pkg/account"
	"transactor-server/pkg/config"
	"transactor-server/pkg/pkgerr"
	"transactor-server/pkg/transaction"

	zapotlp "github.com/SigNoz/zap_otlp"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"go.uber.org/zap"
)

// NewRouter returns a new fiber app with transaction and account api routed on /api/v1/
// it starts a new tracing request if none is present in incoming http request
// it adds a logging middleware which has trace_id and span_id for correlation
// it adds a healthpoint middleware
// it adds a handler to show swagger UI
// and setups up auth for the /api/v1 routes
// @title Transactions Service
// @version 1.0
// @description This is a server which store accounts and transaction details

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description					A Basic way to secure APIs
func NewRouter(
	apiKey string,
	transactionAPI *transaction.API,
	accountAPI *account.API,

	logger *zap.Logger,
) *fiber.App {
	// create a new fiber app
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		// create a custom error handler
		ErrorHandler: ErrorHandler,
	})

	app.Use(recover.New())
	app.Use(healthcheck.New()) // this makes sure our container is recognized as healthy

	app.Get("/swagger/*", swagger.HandlerDefault) // show swagger ui

	apiRouter := app.Group("/api/v1",
		// this mount open telemetry middleware
		// this is responsible for creating and propogating tracing request for http call
		// this middleware also send a set of standard http metrics -
		// http.server.duration
		// http.server.request.size
		// http.server.response.size
		// http.server.active_requests
		otelfiber.Middleware(otelfiber.WithServerName(config.AppName)),

		// this setups a loggig middleware which logs an entry at the end of each request indicating the status, method, etc.
		// it also makes sure to log trace_id and span_id to the log for correlation with a trace :)
		fiberzap.New(fiberzap.Config{
			Logger: logger.With(zap.String("layer", "transport")),
			FieldsFunc: func(c *fiber.Ctx) []zap.Field {
				return []zap.Field{
					zapotlp.SpanCtx(c.UserContext()), // this extracts the span details and creates 2 fields span_id & trace_id
				}
			},
		}),

		// this middle does a basic ApiKey based authentication
		// it tries to match the provided APIKey in the Authorization header
		keyauth.New(keyauth.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return pkgerr.NewServiceError("auth", "", http.StatusForbidden, err.Error())
			},
			KeyLookup: "header:authorization",
			Validator: func(c *fiber.Ctx, s string) (bool, error) {
				if s == apiKey {
					return true, nil
				}
				return false, fmt.Errorf("error validating api key")
			},
		}),
	)

	// mount transaction api routes on /api/v1/transactions
	transactionAPI.Handle(apiRouter.Group("/transactions"))
	// mount account api routes on /api/v1/accounts
	accountAPI.Handle(apiRouter.Group("/accounts"))

	return app
}
