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

func TestDAOBalance(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	client.OperationType.Create().SetDescription("debit").SetID(1).SetIsDebit(true).ExecX(context.Background())
	client.OperationType.Create().SetDescription("credit").SetID(4).SetIsDebit(false).ExecX(context.Background())

	client.Account.Create().SetDocumentNumber("12345").SetID(1).SetName("John Doe").ExecX(context.Background())

	dao := transaction.NewDAO(client)

	ctx := context.Background()

	_, err := dao.Create(ctx, &transaction.CreateRequest{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          -50,
	})
	require.NoError(t, err)

	_, err = dao.Create(ctx, &transaction.CreateRequest{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          -23.5,
	})
	require.NoError(t, err)

	_, err = dao.Create(ctx, &transaction.CreateRequest{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          -18.7,
	})
	require.NoError(t, err)

	firstCreditTxn, err := dao.Create(ctx, &transaction.CreateRequest{
		AccountID:       1,
		OperationTypeID: 4,
		Amount:          60,
	})
	require.NoError(t, err)

	require.Equal(t, 0., firstCreditTxn.Balance)

	firstTxn, err := client.Transaction.Get(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, 0., firstTxn.Balance)

	secondTxn, err := client.Transaction.Get(ctx, 2)
	require.NoError(t, err)
	require.Equal(t, -13.5, secondTxn.Balance)

	thirdTxn, err := client.Transaction.Get(ctx, 3)
	require.NoError(t, err)
	require.Equal(t, -18.7, thirdTxn.Balance)

	secondCreditTxn, err := dao.Create(ctx, &transaction.CreateRequest{
		AccountID:       1,
		OperationTypeID: 4,
		Amount:          100,
	})
	require.NoError(t, err)

	require.Equal(t, 67.8, secondCreditTxn.Balance)

	firstTxn, err = client.Transaction.Get(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, 0., firstTxn.Balance)

	secondTxn, err = client.Transaction.Get(ctx, 2)
	require.NoError(t, err)
	require.Equal(t, 0., secondTxn.Balance)

	thirdTxn, err = client.Transaction.Get(ctx, 3)
	require.NoError(t, err)
	require.Equal(t, 0., thirdTxn.Balance)
}
