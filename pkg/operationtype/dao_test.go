package operationtype_test

import (
	"context"
	"testing"
	"transactor-server/pkg/db/ent/enttest"
	"transactor-server/pkg/operationtype"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestDAOGet(t *testing.T) {
	t.Parallel()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	client.OperationType.Create().SetDescription("debit").SetID(1).SetIsDebit(true).ExecX(context.Background())
	client.OperationType.Create().SetDescription("credit").SetID(2).SetIsDebit(false).ExecX(context.Background())
	client.OperationType.Create().SetDescription("other debit").SetID(3).SetIsDebit(true).ExecX(context.Background())

	dao := operationtype.NewDAO(client)

	resp, err := dao.Get(context.Background(), 2)

	require.NoError(t, err)
	require.NotNil(t, resp)

	require.Equal(t, 2, resp.ID)
	require.Equal(t, "credit", resp.Description)
	require.Equal(t, false, resp.IsDebit)

	resp, err = dao.Get(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)

	require.Equal(t, 1, resp.ID)
	require.Equal(t, "debit", resp.Description)
	require.Equal(t, true, resp.IsDebit)
}
