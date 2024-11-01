package account_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
	"transactor-server/pkg/account"
	"transactor-server/pkg/db/ent"
	"transactor-server/pkg/mocks"
	"transactor-server/pkg/pkgerr"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestServiceCreate(t *testing.T) {
	t.Run("validation errors", func(t *testing.T) {
		t.Parallel()
		service := account.NewService(mocks.NewMockAccountDAO(t), zap.NewNop())

		resp, err := service.Create(context.Background(), &account.CreateRequest{})

		require.Error(t, err)
		require.Nil(t, resp)
		validationErr, ok := err.(*pkgerr.ValidationError)
		require.True(t, ok)
		require.Equal(t, http.StatusBadRequest, validationErr.HttpStatusCode())
		require.NotNil(t, validationErr.ResponseBody())
	})

	t.Run("create db error", func(t *testing.T) {
		t.Parallel()
		accountDAO := mocks.NewMockAccountDAO(t)

		service := account.NewService(accountDAO, zap.NewNop())

		accountDAO.On("Create", mock.Anything, &account.CreateRequest{
			DocumentNumber: "12345",
			Name:           "John Doe",
		}).Return(nil, &ent.ConstraintError{})

		resp, err := service.Create(context.Background(), &account.CreateRequest{
			DocumentNumber: "12345",
			Name:           "John Doe",
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
		accountDAO := mocks.NewMockAccountDAO(t)

		service := account.NewService(accountDAO, zap.NewNop())

		accountDAO.On("Create", mock.Anything, &account.CreateRequest{
			DocumentNumber: "12345",
			Name:           "John Doe",
		}).Return(&ent.Account{ID: 1}, nil)

		resp, err := service.Create(context.Background(), &account.CreateRequest{
			DocumentNumber: "12345",
			Name:           "John Doe",
		})

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 1, resp.ID)
	})
}

func TestServiceGet(t *testing.T) {
	t.Run("validation error", func(t *testing.T) {
		t.Parallel()
		accountDAO := mocks.NewMockAccountDAO(t)

		service := account.NewService(accountDAO, zap.NewNop())

		resp, err := service.Get(context.Background(), -1)

		require.Error(t, err)
		require.Nil(t, resp)
		validationErr, ok := err.(*pkgerr.ValidationError)
		require.True(t, ok)
		require.Equal(t, http.StatusBadRequest, validationErr.HttpStatusCode())
		require.NotNil(t, validationErr.ResponseBody())
	})

	t.Run("db not found error", func(t *testing.T) {
		t.Parallel()
		accountDAO := mocks.NewMockAccountDAO(t)

		service := account.NewService(accountDAO, zap.NewNop())

		accountDAO.On("Get", mock.Anything, 373).Return(nil, &ent.NotFoundError{})

		resp, err := service.Get(context.Background(), 373)

		require.Error(t, err)
		require.Nil(t, resp)
		validationErr, ok := err.(*pkgerr.ServiceError)
		require.True(t, ok)
		require.Equal(t, http.StatusNotFound, validationErr.HttpStatusCode())
		require.NotNil(t, validationErr.ResponseBody())
	})

	t.Run("db some other error", func(t *testing.T) {
		t.Parallel()
		accountDAO := mocks.NewMockAccountDAO(t)

		service := account.NewService(accountDAO, zap.NewNop())

		accountDAO.On("Get", mock.Anything, 373).Return(nil, fmt.Errorf("some error"))

		resp, err := service.Get(context.Background(), 373)

		require.Error(t, err)
		require.Nil(t, resp)
		validationErr, ok := err.(*pkgerr.ServiceError)
		require.True(t, ok)
		require.Equal(t, http.StatusInternalServerError, validationErr.HttpStatusCode())
		require.NotNil(t, validationErr.ResponseBody())
	})

	t.Run("no error", func(t *testing.T) {
		t.Parallel()
		accountDAO := mocks.NewMockAccountDAO(t)

		service := account.NewService(accountDAO, zap.NewNop())

		dbAccount := &ent.Account{
			ID:             373,
			CreateTime:     time.Now().AddDate(-1, 0, 0),
			UpdateTime:     time.Now(),
			Name:           "John Doe",
			DocumentNumber: "12345",
		}

		accountDAO.
			On("Get", mock.Anything, 373).
			Return(dbAccount, nil)

		resp, err := service.Get(context.Background(), 373)

		require.NoError(t, err)
		require.NotNil(t, resp)

		require.Equal(t, resp.ID, dbAccount.ID)
		require.Equal(t, resp.Name, dbAccount.Name)
		require.Equal(t, resp.DocumentNumber, dbAccount.DocumentNumber)
		require.Equal(t, resp.CreatedAt, dbAccount.CreateTime)
		require.Equal(t, resp.UpdatedAt, dbAccount.UpdateTime)
	})
}
