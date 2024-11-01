package account

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateRequest struct {
	DocumentNumber string `json:"document_number"`
	Name           string `json:"name"`
}

// Validate validates the CreateRequest to
// have non empty document_number and
// have name to be >= 8 & <= 100 characters in length
func (req CreateRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.DocumentNumber, validation.Required),
		validation.Field(&req.Name, validation.Required, validation.Length(8, 100)),
	)
}

type CreateResponse struct {
	ID int `json:"id"`
}

type Account struct {
	ID             int       `json:"id"`
	DocumentNumber string    `json:"document_number"`
	Name           string    `json:"name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
