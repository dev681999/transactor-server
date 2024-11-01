package operationtype

import (
	"context"
	"transactor-server/pkg/db/ent"
)

// DAO defines the data access object interface for operation_type model
//
//go:generate go run -mod=mod github.com/vektra/mockery/v2 --name DAO --output ../mocks --structname MockOperationTypeDAO --filename operationtype_dao.go
type DAO interface {
	// Get tries to find an existing operation type in DB by id
	Get(ctx context.Context, id int) (*ent.OperationType, error)
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

func (d *dao) Get(ctx context.Context, id int) (*ent.OperationType, error) {
	return d.entClient.OperationType.Get(ctx, id)
}
