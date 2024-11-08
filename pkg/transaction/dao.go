package transaction

import (
	"context"
	"time"
	"transactor-server/pkg/db/ent"
	"transactor-server/pkg/db/ent/transaction"

	"entgo.io/ent/dialect/sql"
)

// DAO defines the data access object interface for transaction model
//
//go:generate go run -mod=mod github.com/vektra/mockery/v2 --name DAO --output ../mocks --structname MockTransactionDAO --filename transaction_dao.go
type DAO interface {
	Create(ctx context.Context, req *CreateRequest) (*ent.Transaction, error)
}

type dao struct {
	entClient *ent.Client
}

var _ DAO = (*dao)(nil)

// NewDAO returns a new DAO which use ent as database orm
func NewDAO(entClient *ent.Client) DAO {
	return &dao{
		entClient: entClient,
	}
}

func (d *dao) Create(ctx context.Context, req *CreateRequest) (*ent.Transaction, error) {
	tx, err := d.entClient.Tx(ctx)
	if err != nil {
		return nil, err
	}

	balance := req.Amount

	lastID := 0

	for balance > 0 {
		balanceTransactions, err := tx.Transaction.
			Query().
			Where(
				transaction.BalanceLT(0),
				transaction.AccountID(req.AccountID),
				transaction.IDGT(lastID),
			).
			Order(
				transaction.ByID(sql.OrderAsc()),
			).
			Limit(10).
			All(ctx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		if len(balanceTransactions) < 1 {
			break
		}

		for _, balanceTransaction := range balanceTransactions {
			txBalance := balanceTransaction.Balance * -1
			if txBalance >= balance {
				err = tx.Transaction.
					UpdateOneID(balanceTransaction.ID).
					SetBalance(balanceTransaction.Balance + balance).
					Exec(ctx)
				if err != nil {
					tx.Rollback()
					return nil, err
				}

				balance = 0
			} else if balance > txBalance {
				err = tx.Transaction.
					UpdateOneID(balanceTransaction.ID).
					SetBalance(0).
					Exec(ctx)
				if err != nil {
					tx.Rollback()
					return nil, err
				}
				balance -= txBalance
			}

			if balance == 0 {
				break
			}
		}

		lastID = balanceTransactions[len(balanceTransactions)-1].ID
	}

	dbTxn, err := tx.Transaction.
		Create().
		SetAccountID(req.AccountID).
		SetOperationTypeID(req.OperationTypeID).
		SetTimestamp(time.Now()).
		SetAmount(req.Amount).
		SetBalance(balance).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return dbTxn, nil
}
