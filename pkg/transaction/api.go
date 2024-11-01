package transaction

import (
	"net/http"

	"transactor-server/pkg/pkgerr"

	"github.com/gofiber/fiber/v2"
)

// API is the api handler for transaction apis
type API struct {
	sevice Service
}

// NewAPI returns a new API handler ready to handle routes
func NewAPI(service Service) *API {
	return &API{
		sevice: service,
	}
}

// Handle sets up all the routes with their handler funcs for transaction apis
func (a *API) Handle(router fiber.Router) {
	router.Post("/", a.createTransaction)
}

// createTransaction creates a new transaction in DB
// @Summary      create a transaction
// @Produce      json
// @Tags		 transaction
// @Param        req    body     CreateRequest  true  "transaction details to create"
// @Success      201  {object}  CreateResponse
// @Failure      400  {object}  pkgerr.ValidationErrorResponseBody
// @Failure      500  {object}  pkgerr.ServiceErrorResponseBody
// @Security	 ApiKeyAuth
// @Router       /api/v1/transactions [post]
func (a *API) createTransaction(c *fiber.Ctx) error {
	req := &CreateRequest{}

	// try to parse the body
	err := c.BodyParser(req)
	if err != nil {
		return pkgerr.NewServiceError("account", "body_parse_failure", http.StatusBadRequest, err.Error())
	}

	// call the sevice to create the transaction
	// UserContext returns the actual context will all opentelemetry related details added
	resp, err := a.sevice.Create(c.UserContext(), req)
	if err != nil {
		return err
	}

	// incase of no error we return the response with 201 status
	return c.Status(http.StatusCreated).JSON(resp)
}
