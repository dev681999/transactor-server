package account

import (
	"context"
	"transactor-server/pkg/db/ent"
)

// DAO defines the data access object interface for account model
//
//go:generate go run -mod=mod github.com/vektra/mockery/v2 --name DAO --output ../mocks --structname MockAccountDAO --filename account_dao.go
type DAO interface {
	// Create inserts a new account record in DB
	Create(ctx context.Context, req *CreateRequest) (*ent.Account, error)
	// Get tries to find an existing account record in DB by id
	Get(ctx context.Context, id int) (*ent.Account, error)
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

func (d *dao) Create(ctx context.Context, req *CreateRequest) (*ent.Account, error) {
	return d.entClient.Account.
		Create().
		SetDocumentNumber(req.DocumentNumber).
		SetName(req.Name).
		Save(ctx)
}

func (d *dao) Get(ctx context.Context, id int) (*ent.Account, error) {
	return d.entClient.Account.Get(ctx, id)
}
