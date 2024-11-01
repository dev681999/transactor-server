package transaction_test

import (
	"context"
	"net/http"
	"testing"
	"transactor-server/pkg/db/ent"
	"transactor-server/pkg/mocks"
	"transactor-server/pkg/pkgerr"
	"transactor-server/pkg/transaction"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestServiceCreate(t *testing.T) {
	t.Run("validation errors", func(t *testing.T) {
		t.Parallel()
		service := transaction.NewService(mocks.NewMockOperationTypeDAO(t), mocks.NewMockTransactionDAO(t), zap.NewNop())

		resp, err := service.Create(context.Background(), &transaction.CreateRequest{})

		require.Error(t, err)
		require.Nil(t, resp)
		validationErr, ok := err.(*pkgerr.ValidationError)
		require.True(t, ok)
		require.Equal(t, http.StatusBadRequest, validationErr.HttpStatusCode())
		require.NotNil(t, validationErr.ResponseBody())
	})

	t.Run("operation type not found", func(t *testing.T) {
		t.Parallel()
		operationTypeDAO := mocks.NewMockOperationTypeDAO(t)
		transactionDAO := mocks.NewMockTransactionDAO(t)

		service := transaction.NewService(operationTypeDAO, transactionDAO, zap.NewNop())

		operationTypeDAO.On("Get", mock.Anything, 1).Return(nil, &ent.NotFoundError{})

		resp, err := service.Create(context.Background(), &transaction.CreateRequest{
			AccountID:       1,
			OperationTypeID: 1,
			Amount:          98.99,
		})

		require.Error(t, err)
		require.Nil(t, resp)
		serviceErr, ok := err.(*pkgerr.ServiceError)
		require.True(t, ok)
		require.Equal(t, http.StatusNotFound, serviceErr.HttpStatusCode())
		require.NotNil(t, serviceErr.ResponseBody())
	})

	t.Run("operation type amount sign mismatch for debit", func(t *testing.T) {
		t.Parallel()
		operationTypeDAO := mocks.NewMockOperationTypeDAO(t)
		transactionDAO := mocks.NewMockTransactionDAO(t)

		service := transaction.NewService(operationTypeDAO, transactionDAO, zap.NewNop())

		operationTypeDAO.On("Get", mock.Anything, 1).Return(&ent.OperationType{
			ID:      1,
			IsDebit: true,
		}, nil)

		resp, err := service.Create(context.Background(), &transaction.CreateRequest{
			AccountID:       1,
			OperationTypeID: 1,
			Amount:          98.99,
		})

		require.Error(t, err)
		require.Nil(t, resp)
		serviceErr, ok := err.(*pkgerr.ServiceError)
		require.True(t, ok)
		require.Equal(t, http.StatusBadRequest, serviceErr.HttpStatusCode())
		require.NotNil(t, serviceErr.ResponseBody())
	})

	t.Run("operation type amount sign mismatch for credit", func(t *testing.T) {
		t.Parallel()
		operationTypeDAO := mocks.NewMockOperationTypeDAO(t)
		transactionDAO := mocks.NewMockTransactionDAO(t)

		service := transaction.NewService(operationTypeDAO, transactionDAO, zap.NewNop())

		operationTypeDAO.On("Get", mock.Anything, 1).Return(&ent.OperationType{
			ID:      1,
			IsDebit: false,
		}, nil)

		resp, err := service.Create(context.Background(), &transaction.CreateRequest{
			AccountID:       1,
			OperationTypeID: 1,
			Amount:          -98.99,
		})

		require.Error(t, err)
		require.Nil(t, resp)
		serviceErr, ok := err.(*pkgerr.ServiceError)
		require.True(t, ok)
		require.Equal(t, http.StatusBadRequest, serviceErr.HttpStatusCode())
		require.NotNil(t, serviceErr.ResponseBody())
	})

	t.Run("create db error", func(t *testing.T) {
		t.Parallel()
		operationTypeDAO := mocks.NewMockOperationTypeDAO(t)
		transactionDAO := mocks.NewMockTransactionDAO(t)

		service := transaction.NewService(operationTypeDAO, transactionDAO, zap.NewNop())

		operationTypeDAO.On("Get", mock.Anything, 1).Return(&ent.OperationType{
			ID:      1,
			IsDebit: false,
		}, nil)

		transactionDAO.On("Create", mock.Anything, &transaction.CreateRequest{
			AccountID:       1,
			OperationTypeID: 1,
			Amount:          98.99,
		}).Return(nil, &ent.ConstraintError{})

		resp, err := service.Create(context.Background(), &transaction.CreateRequest{
			AccountID:       1,
			OperationTypeID: 1,
			Amount:          98.99,
		})

		require.Error(t, err)
		require.Nil(t, resp)
		serviceErr, ok := err.(*pkgerr.ServiceError)
		require.True(t, ok)
		require.Equal(t, http.StatusBadRequest, serviceErr.HttpStatusCode())
		require.NotNil(t, serviceErr.ResponseBody())
	})

	t.Run("no error", func(t *testing.T) {
		t.Parallel()
		operationTypeDAO := mocks.NewMockOperationTypeDAO(t)
		transactionDAO := mocks.NewMockTransactionDAO(t)

		service := transaction.NewService(operationTypeDAO, transactionDAO, zap.NewNop())

		operationTypeDAO.On("Get", mock.Anything, 1).Return(&ent.OperationType{
			ID:      1,
			IsDebit: false,
		}, nil)

		transactionDAO.On("Create", mock.Anything, &transaction.CreateRequest{
			AccountID:       1,
			OperationTypeID: 1,
			Amount:          98.99,
		}).Return(&ent.Transaction{ID: 1}, nil)

		resp, err := service.Create(context.Background(), &transaction.CreateRequest{
			AccountID:       1,
			OperationTypeID: 1,
			Amount:          98.99,
		})

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 1, resp.ID)
	})
}
