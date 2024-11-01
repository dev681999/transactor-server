package transaction

import (
	"context"
	"time"
	"transactor-server/pkg/db/ent"
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
	return d.entClient.Transaction.
		Create().
		SetAccountID(req.AccountID).
		SetOperationTypeID(req.OperationTypeID).
		SetTimestamp(time.Now()).
		SetAmount(req.Amount).
		Save(ctx)
}
