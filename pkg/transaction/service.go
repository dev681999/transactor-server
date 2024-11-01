package transaction

import (
	"context"
	"net/http"
	"transactor-server/pkg/config"
	"transactor-server/pkg/infra/log"
	"transactor-server/pkg/operationtype"
	"transactor-server/pkg/pkgerr"

	zapotlp "github.com/SigNoz/zap_otlp"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"
)

// Service handles the main business logic for transaction related things
//
//go:generate go run -mod=mod github.com/vektra/mockery/v2 --name Service --output ../mocks --structname MockTransactionService  --filename transaction_service.go
type Service interface {
	// Create creates a new transaction record in the database layer
	Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error)
}

// a traced, logged and metered transaction service
type service struct {
	operationtypeDAO operationtype.DAO
	transactionDAO   DAO

	createCounterSuccess metric.Int64Counter
	createCounterFailure metric.Int64Counter

	logger *zap.Logger
}

var _ Service = (*service)(nil)

func NewService(
	operationtypeDAO operationtype.DAO,
	transactionDAO DAO,

	logger *zap.Logger,
) Service {
	meter := otel.GetMeterProvider().Meter("transactor-server")

	createCounterSuccess, err := meter.Int64Counter("transaction_service_create_success")
	if err != nil {
		log.L.Fatal("", zap.Error(err))
	}
	createCounterFailure, err := meter.Int64Counter("transaction_service_create_failure")
	if err != nil {
		log.L.Fatal("", zap.Error(err))
	}

	return &service{
		operationtypeDAO:     operationtypeDAO,
		transactionDAO:       transactionDAO,
		createCounterSuccess: createCounterSuccess,
		createCounterFailure: createCounterFailure,

		logger: logger,
	}
}

var (
	// ErrOperationTypeAmountSignMismatch indicates the amount sign for specified operatio type is not correct
	ErrOperationTypeAmountSignMismatch = pkgerr.NewServiceError(
		"transaction", "operator_type_amount_sign_mismatch",
		http.StatusBadRequest,
		"operator type amount sign mismatch",
	)
)

func (s *service) Create(ctx context.Context, req *CreateRequest) (resp *CreateResponse, err error) {
	// start span
	ctx, span := otel.Tracer(config.AppName).Start(ctx, "TransactionService.Create")
	// end span before returning
	defer span.End()
	defer func() {
		// incase of error set the span status to error
		if err != nil {
			span.SetStatus(codes.Error, "error")
			span.RecordError(err)
			s.logger.Error("end TransactionService.Create", zapotlp.SpanCtx(ctx), zap.Any("req", req), zap.Error(err))
			s.createCounterFailure.Add(ctx, 1)
		} else {
			s.logger.Info("end TransactionService.Create", zapotlp.SpanCtx(ctx), zap.Any("req", req), zap.Any("resp", resp))
			s.createCounterSuccess.Add(ctx, 1)
		}
	}()

	s.logger.Info("calling TransactionService.Create", zapotlp.SpanCtx(ctx), zap.Any("req", req))

	// run validations, please the function to know more!
	err = req.Validate()
	if err != nil {
		err = pkgerr.WrapStructValidationError(err)
		return
	}

	// next we try to find the operation type from id sent
	operationtype, err := s.operationtypeDAO.Get(ctx, req.OperationTypeID)
	if err != nil {
		err = pkgerr.WrapDAOError(err)
		return
	}

	// next we want to make sure the sign on amount matches operation type
	if operationtype.IsDebit && req.Amount > 0 {
		err = ErrOperationTypeAmountSignMismatch
		return
	} else if !operationtype.IsDebit && req.Amount < 0 {
		err = ErrOperationTypeAmountSignMismatch
		return
	}

	// fianll call dao to insert the record in db
	dbTransaction, err := s.transactionDAO.Create(ctx, req)
	if err != nil {
		err = pkgerr.WrapDAOError(err)
		return
	}

	// send back id of newly created transction
	return &CreateResponse{
		ID: dbTransaction.ID,
	}, nil
}
