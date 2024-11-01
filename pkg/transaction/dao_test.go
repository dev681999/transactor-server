package transaction_test

import (
	"context"
	"testing"
	"transactor-server/pkg/db/ent/enttest"
	"transactor-server/pkg/transaction"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestDAOCreate(t *testing.T) {
	t.Parallel()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	client.OperationType.Create().SetDescription("debit").SetID(1).SetIsDebit(true).ExecX(context.Background())
	client.Account.Create().SetDocumentNumber("12345").SetID(373).SetName("John Doe").ExecX(context.Background())

	dao := transaction.NewDAO(client)

	resp, err := dao.Create(context.Background(), &transaction.CreateRequest{
		AccountID:       373,
		OperationTypeID: 1,
		Amount:          -39.88,
	})

	require.NoError(t, err)
	require.NotNil(t, resp)

	require.Equal(t, 373, resp.AccountID)
	require.Equal(t, 1, resp.OperationTypeID)
	require.Equal(t, -39.88, resp.Amount)
	require.Equal(t, 1, resp.ID)

	dbResp := client.Transaction.Query().OnlyX(context.Background())
	require.Equal(t, 373, dbResp.AccountID)
	require.Equal(t, 1, dbResp.OperationTypeID)
	require.Equal(t, -39.88, dbResp.Amount)
	require.Equal(t, 1, dbResp.ID)
}
