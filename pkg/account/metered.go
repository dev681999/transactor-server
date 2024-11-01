package account

import (
	"context"
	"transactor-server/pkg/infra/log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"
)

// meteredSevice is a middleware/wrapper to the account.Service
// it like name suggests adds metrics to each service call
// on success it adds account_service_get_success and on failure account_service_get_failure
type meteredSevice struct {
	service Service

	meter metric.Meter

	getCounterSuccess metric.Int64Counter
	getCounterFailure metric.Int64Counter

	createCounterSuccess metric.Int64Counter
	createCounterFailure metric.Int64Counter
}

var _ Service = (*meteredSevice)(nil)

func NewMeteredService(service Service) Service {
	meter := otel.GetMeterProvider().Meter("transactor-server")

	getCounterSuccess, err := meter.Int64Counter("account_service_get_success")
	if err != nil {
		log.L.Fatal("", zap.Error(err))
	}
	getCounterFailure, err := meter.Int64Counter("account_service_get_failure")
	if err != nil {
		log.L.Fatal("", zap.Error(err))
	}

	createCounterSuccess, err := meter.Int64Counter("account_service_create_success")
	if err != nil {
		log.L.Fatal("", zap.Error(err))
	}
	createCounterFailure, err := meter.Int64Counter("account_service_create_failure")
	if err != nil {
		log.L.Fatal("", zap.Error(err))
	}

	return &meteredSevice{
		service:              service,
		meter:                meter,
		getCounterSuccess:    getCounterSuccess,
		getCounterFailure:    getCounterFailure,
		createCounterSuccess: createCounterSuccess,
		createCounterFailure: createCounterFailure,
	}
}

func (m *meteredSevice) Create(ctx context.Context, req *CreateRequest) (resp *CreateResponse, err error) {
	defer func() {
		if err == nil {
			m.createCounterSuccess.Add(ctx, 1)
		} else {
			m.createCounterFailure.Add(ctx, 1)
		}
	}()
	resp, err = m.service.Create(ctx, req)
	return
}

func (m *meteredSevice) Get(ctx context.Context, id int) (resp *Account, err error) {
	defer func() {
		if err == nil {
			m.getCounterSuccess.Add(ctx, 1)
		} else {
			m.getCounterFailure.Add(ctx, 1)
		}
	}()
	resp, err = m.service.Get(ctx, id)
	return
}
