package account_test

import (
	"context"
	"testing"
	"transactor-server/pkg/account"
	"transactor-server/pkg/db/ent/enttest"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestDAOCreate(t *testing.T) {
	t.Parallel()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	dao := account.NewDAO(client)

	resp, err := dao.Create(context.Background(), &account.CreateRequest{
		DocumentNumber: "12345",
		Name:           "John Doe",
	})

	require.NoError(t, err)
	require.NotNil(t, resp)

	require.Equal(t, "12345", resp.DocumentNumber)
	require.Equal(t, "John Doe", resp.Name)
	require.Equal(t, 1, resp.ID)

	dbResp := client.Account.Query().OnlyX(context.Background())
	require.Equal(t, "12345", dbResp.DocumentNumber)
	require.Equal(t, "John Doe", dbResp.Name)
	require.Equal(t, 1, dbResp.ID)
}

func TestDAOGet(t *testing.T) {
	t.Parallel()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	dao := account.NewDAO(client)

	_, err := dao.Create(context.Background(), &account.CreateRequest{
		DocumentNumber: "12345",
		Name:           "John Doe",
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = dao.Create(context.Background(), &account.CreateRequest{
		DocumentNumber: "78901",
		Name:           "Jane Doe",
	})
	if err != nil {
		t.Fatal(err)
	}

	resp, err := dao.Get(context.Background(), 2)

	require.NoError(t, err)
	require.NotNil(t, resp)

	require.Equal(t, 2, resp.ID)
	require.Equal(t, "78901", resp.DocumentNumber)
	require.Equal(t, "Jane Doe", resp.Name)

	resp, err = dao.Get(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)

	require.Equal(t, 1, resp.ID)
	require.Equal(t, "12345", resp.DocumentNumber)
	require.Equal(t, "John Doe", resp.Name)
}
