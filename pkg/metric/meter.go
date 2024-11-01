package metric

import (
	"context"
	"transactor-server/pkg/infra/log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	metricsdk "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

// Init setups the global open telemetry metric provider which connects to a open telemetry colector on grpc
func Init(collectorURL string, res *resource.Resource) func(context.Context) error {
	exporter, err := otlpmetricgrpc.New(
		context.Background(),
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint(collectorURL),
	)

	if err != nil {
		log.L.Sugar().Fatalf("Failed to create exporter: %v", err)
	}

	otel.SetMeterProvider(
		metricsdk.NewMeterProvider(
			metricsdk.WithReader(metricsdk.NewPeriodicReader(exporter)),
			metricsdk.WithResource(res),
		),
	)

	return exporter.Shutdown
}
