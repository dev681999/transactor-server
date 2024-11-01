package account_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"transactor-server/pkg/account"
	"transactor-server/pkg/api"
	"transactor-server/pkg/mocks"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

var setupApp = func(t *testing.T) (*fiber.App, *mocks.MockAccountService) {
	app := fiber.New(fiber.Config{
		ErrorHandler: api.ErrorHandler,
	})

	router := app.Group("/test/accounts")
	service := mocks.NewMockAccountService(t)

	api := account.NewAPI(service)
	api.Handle(router)

	return app, service
}

func TestAPICreate(t *testing.T) {
	t.Run("body parsing errpr", func(t *testing.T) {
		t.Parallel()
		app, _ := setupApp(t)

		req := httptest.NewRequest(http.MethodPost, "/test/accounts/", bytes.NewBufferString("something"))

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.NotNil(t, resp)

		require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("service errpr", func(t *testing.T) {
		t.Parallel()
		app, service := setupApp(t)

		body := &account.CreateRequest{
			DocumentNumber: "12345",
			Name:           "John Doe",
		}

		b, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/test/accounts/", bytes.NewBuffer(b))
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

		body := &account.CreateRequest{
			DocumentNumber: "12345",
			Name:           "John Doe",
		}

		b, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/test/accounts/", bytes.NewBuffer(b))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		service.On("Create", mock.Anything, body).Return(&account.CreateResponse{
			ID: 373,
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

		require.Equal(t, int64(373), gjson.Get(string(b), "id").Int())
	})
}

func TestAPIGet(t *testing.T) {
	t.Run("invalid id", func(t *testing.T) {
		t.Parallel()
		app, _ := setupApp(t)

		req := httptest.NewRequest(http.MethodGet, "/test/accounts/abc", nil)

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.NotNil(t, resp)

		require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("service error", func(t *testing.T) {
		t.Parallel()
		app, service := setupApp(t)

		req := httptest.NewRequest(http.MethodGet, "/test/accounts/373", nil)
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		service.On("Get", mock.Anything, 373).Return(nil, fmt.Errorf("some error"))

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.NotNil(t, resp)

		require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("no error", func(t *testing.T) {
		t.Parallel()
		app, service := setupApp(t)

		req := httptest.NewRequest(http.MethodGet, "/test/accounts/373", nil)
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		acc := &account.Account{
			ID:             373,
			DocumentNumber: "12345",
			Name:           "John Doe",
			CreatedAt:      time.Now().Add(time.Hour * -24).Truncate(time.Millisecond),
			UpdatedAt:      time.Now().Truncate(time.Millisecond),
		}

		service.On("Get", mock.Anything, 373).Return(acc, nil)

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.NotNil(t, resp)

		require.Equal(t, http.StatusOK, resp.StatusCode)

		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		require.Equal(t, int64(373), gjson.Get(string(b), "id").Int())
		require.Equal(t, "12345", gjson.Get(string(b), "document_number").String())
		require.Equal(t, "John Doe", gjson.Get(string(b), "name").String())

		createdAtStr := gjson.Get(string(b), "created_at").String()
		returnedCreatedAt, err := time.Parse(time.RFC3339, createdAtStr)
		if err != nil {
			t.Fatal(err)
		}

		require.Equal(t, acc.CreatedAt.UTC(), returnedCreatedAt.UTC())

		updatedAtStr := gjson.Get(string(b), "updated_at").String()
		returnedUpdatedAt, err := time.Parse(time.RFC3339, updatedAtStr)
		if err != nil {
			t.Fatal(err)
		}

		require.Equal(t, acc.UpdatedAt.UTC(), returnedUpdatedAt.UTC())
	})
}
