package account

import "transactor-server/pkg/db/ent"

// MapEntAccountToAccount maps an ent.Account record to account.Account model
func MapEntAccountToAccount(a *ent.Account) *Account {
	if a == nil {
		return nil
	}

	return &Account{
		ID:             a.ID,
		DocumentNumber: a.DocumentNumber,
		Name:           a.Name,
		CreatedAt:      a.CreateTime,
		UpdatedAt:      a.UpdateTime,
	}
}
