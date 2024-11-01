package transaction

import validation "github.com/go-ozzo/ozzo-validation/v4"

type CreateRequest struct {
	AccountID       int     `json:"account_id"`
	OperationTypeID int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

// Validate validates the CreateRequest to
// have +ve account and operation type id and
// ensure amount is not ZERO
func (req CreateRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.AccountID, validation.Min(1)),
		validation.Field(&req.OperationTypeID, validation.Min(1)),
		validation.Field(&req.Amount, validation.Required), // if there can be 0 amount then we should remove this!
	)
}

type CreateResponse struct {
	ID int `json:"id"`
}
