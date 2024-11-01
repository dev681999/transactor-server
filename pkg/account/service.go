package account

import (
	"context"
	"transactor-server/pkg/pkgerr"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.uber.org/zap"
)

// Service handles the main business logic for account related things
//
//go:generate go run -mod=mod github.com/vektra/mockery/v2 --name Service --output ../mocks --structname MockAccountService  --filename account_service.go
type Service interface {
	// Create creates a new account record in the database layer
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Get tries to find an existing account in dtabase layer
	Get(context.Context, int) (*Account, error)
}

type service struct {
	accountDAO DAO

	logger *zap.Logger
}

var _ Service = (*service)(nil)

func NewService(
	accountDAO DAO,

	logger *zap.Logger,
) Service {
	return &service{
		accountDAO: accountDAO,

		logger: logger,
	}
}

func (s *service) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	// run validations, please the function to know more!
	err := req.Validate()
	if err != nil {
		return nil, pkgerr.WrapStructValidationError(err)
	}

	// calls dao to insert record in database
	dbAccount, err := s.accountDAO.Create(ctx, req)
	if err != nil {
		return nil, pkgerr.WrapDAOError(err)
	}

	// return id of newly created account
	return &CreateResponse{
		ID: dbAccount.ID,
	}, nil
}

func (s *service) Get(ctx context.Context, id int) (*Account, error) {
	// validates the id to be +ve
	err := validation.Validate(id, validation.Min(1))
	if err != nil {
		return nil, pkgerr.WrapValidationError(err, "id")
	}

	// calls dao to get the record from database
	dbAccount, err := s.accountDAO.Get(ctx, id)
	if err != nil {
		return nil, pkgerr.WrapDAOError(err)
	}

	// map the ent model to return format
	account := MapEntAccountToAccount(dbAccount)

	return account, nil
}
