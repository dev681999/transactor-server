package log

import (
	"os"

	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"

	zapotlpencoder "github.com/SigNoz/zap_otlp/zap_otlp_encoder"
	zapotlpsync "github.com/SigNoz/zap_otlp/zap_otlp_sync"
	"github.com/samber/lo"
)

// Logger is default logger
var Logger = zap.L()
var L = zap.L()

// New sets up logger
func New(debug bool, otelConn *grpc.ClientConn, res *resource.Resource) (logger *zap.Logger, otlpSync *zapotlpsync.OtelSyncer) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeDuration = zapcore.StringDurationEncoder
	consoleEncoder := zapcore.NewJSONEncoder(config)
	defaultLogLevel := lo.Ternary(debug, zapcore.DebugLevel, zap.InfoLevel)

	// if a connection is present setup opentelemetry logging module
	if otelConn != nil {
		otlpEncoder := zapotlpencoder.NewOTLPEncoder(config)

		otlpSync = zapotlpsync.NewOtlpSyncer(otelConn, zapotlpsync.Options{
			BatchSize:      2,
			ResourceSchema: semconv.SchemaURL,
			Resource:       res,
		})

		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, os.Stdout, defaultLogLevel),
			zapcore.NewCore(otlpEncoder, otlpSync, defaultLogLevel),
		)
		logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	} else {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		logger, _ = config.Build()
	}

	zap.ReplaceGlobals(logger)

	Logger = logger
	L = logger

	return
}
