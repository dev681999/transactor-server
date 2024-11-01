package account

import (
	"context"
	"transactor-server/pkg/config"

	zapotlp "github.com/SigNoz/zap_otlp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.uber.org/zap"
)

// tracedService is a middleware/wrapper to the account.Service
// it like name suggests adds a span to the each of the methods
// on error it ensures that span is marked with error
// it also logs request and response and error if any
type tracedService struct {
	service Service
	logger  *zap.Logger
}

func NewTracedService(service Service, logger *zap.Logger) Service {
	return &tracedService{
		service: service,
		logger:  logger,
	}
}

var _ Service = (*tracedService)(nil)

func (t *tracedService) Create(ctx context.Context, req *CreateRequest) (resp *CreateResponse, err error) {
	// start span
	ctx, span := otel.Tracer(config.AppName).Start(ctx, "AccountService.Create")
	// end span before returning
	defer span.End()
	defer func() {
		// incase of error set the span status to error
		if err != nil {
			span.SetStatus(codes.Error, "error")
			span.RecordError(err)
		}
	}()

	t.logger.Info("calling AccountService.Create", zapotlp.SpanCtx(ctx), zap.Any("req", req))

	resp, err = t.service.Create(ctx, req)

	if err != nil {
		t.logger.Error("end AccountService.Create with error", zapotlp.SpanCtx(ctx), zap.Any("req", req), zap.Error(err))
	} else {
		t.logger.Info("end AccountService.Create", zapotlp.SpanCtx(ctx), zap.Any("req", req), zap.Any("resp", resp))
	}

	return
}

func (t *tracedService) Get(ctx context.Context, id int) (resp *Account, err error) {
	// start span
	ctx, span := otel.Tracer(config.AppName).Start(ctx, "AccountService.Get")
	// end span before returning
	defer span.End()
	defer func() {
		// incase of error set the span status to error
		if err != nil {
			span.SetStatus(codes.Error, "error")
			span.RecordError(err)
		}
	}()

	t.logger.Info("calling AccountService.Get", zapotlp.SpanCtx(ctx), zap.Int("id", id))

	resp, err = t.service.Get(ctx, id)

	if err != nil {
		t.logger.Error("end AccountService.Get with error", zapotlp.SpanCtx(ctx), zap.Int("id", id), zap.Error(err))
	} else {
		t.logger.Info("end AccountService.Get", zapotlp.SpanCtx(ctx), zap.Int("id", id), zap.Any("resp", resp))
	}

	return
}
