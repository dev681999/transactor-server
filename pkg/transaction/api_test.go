package transaction_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"transactor-server/pkg/api"
	"transactor-server/pkg/mocks"
	"transactor-server/pkg/transaction"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

var setupApp = func(t *testing.T) (*fiber.App, *mocks.MockTransactionService) {
	app := fiber.New(fiber.Config{
		ErrorHandler: api.ErrorHandler,
	})

	router := app.Group("/test/transactions")
	service := mocks.NewMockTransactionService(t)

	api := transaction.NewAPI(service)
	api.Handle(router)

	return app, service
}

func TestAPICreate(t *testing.T) {
	t.Run("body parsing errpr", func(t *testing.T) {
		t.Parallel()
		app, _ := setupApp(t)

		req := httptest.NewRequest(http.MethodPost, "/test/transactions/", bytes.NewBufferString("something"))

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.NotNil(t, resp)

		require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("service errpr", func(t *testing.T) {
		t.Parallel()
		app, service := setupApp(t)

		body := &transaction.CreateRequest{
			AccountID:       373,
			OperationTypeID: 1,
			Amount:          -98.75,
		}

		b, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/test/transactions/", bytes.NewBuffer(b))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		service.On("Create", mock.Anything, body).Return(nil, fmt.Errorf("some error"))

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.NotNil(t, resp)

		require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("no errpr", func(t *testing.T) {
		t.Parallel()
		app, service := setupApp(t)

		body := &transaction.CreateRequest{
			AccountID:       373,
			OperationTypeID: 1,
			Amount:          -98.75,
		}

		b, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/test/transactions/", bytes.NewBuffer(b))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		service.On("Create", mock.Anything, body).Return(&transaction.CreateResponse{
			ID: 999,
		}, nil)

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.NotNil(t, resp)

		require.Equal(t, http.StatusCreated, resp.StatusCode)

		defer resp.Body.Close()
		b, err = io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		require.Equal(t, int64(999), gjson.Get(string(b), "id").Int())
	})
}
