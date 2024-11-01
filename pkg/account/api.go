package account

import (
	"net/http"
	"strconv"

	"transactor-server/pkg/pkgerr"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// API is the api handler for account apis
type API struct {
	sevice Service
}

// NewAPI returns a new API handler ready to handle routes
func NewAPI(service Service) *API {
	return &API{
		sevice: service,
	}
}

// Handle sets up all the routes with their handler funcs for account apis
func (a *API) Handle(router fiber.Router) {
	router.Post("/", a.createAccount)
	router.Get("/:id", a.getAccount)
}

// createAccount creates a new account in DB
// @Summary      create a account
// @Produce      json
// @Tags		 account
// @Param        req    body     CreateRequest  true  "account details to create"
// @Success      201  {object}  CreateResponse
// @Failure      400  {object}  pkgerr.ValidationErrorResponseBody
// @Failure      500  {object}  pkgerr.ServiceErrorResponseBody
// @Security	 ApiKeyAuth
// @Router       /api/v1/accounts [post]
func (a *API) createAccount(c *fiber.Ctx) error {
	req := &CreateRequest{}

	// try to parse the body
	err := c.BodyParser(req)
	if err != nil {
		return pkgerr.NewServiceError("account", "body_parse_failure", http.StatusBadRequest, err.Error())
	}

	// call the sevice to create the user
	// UserContext returns the actual context will all opentelemetry related details added
	resp, err := a.sevice.Create(c.UserContext(), req)
	if err != nil {
		return err
	}

	// incase of no error we return the response with 201 status
	return c.Status(http.StatusCreated).JSON(resp)
}

// getAccount return an existing account detail
// @Summary      get an account
// @Produce      json
// @Tags		 account
// @Param        id    path     int  true  "account id"
// @Success      200  {object}  Account
// @Failure      400  {object}  pkgerr.ValidationErrorResponseBody
// @Failure      404  {object}  pkgerr.ServiceErrorResponseBody
// @Failure      500  {object}  pkgerr.ServiceErrorResponseBody
// @Security	 ApiKeyAuth
// @Router       /api/v1/accounts/{id} [get]
func (a *API) getAccount(c *fiber.Ctx) error {
	idStr := c.Params("id")
	// technically the id will never be empty bcz empty id means a different route altogether
	// but just to be safe :)
	if lo.IsEmpty(idStr) {
		return pkgerr.NewServiceError("validation", "validation_failed", http.StatusBadRequest, "id path variable is required")
	}

	// we try to parse the id to an int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return pkgerr.NewServiceError("validation", "validation_failed", http.StatusBadRequest, "id path variable must be an integer")
	}

	// call the service to get account details
	resp, err := a.sevice.Get(c.UserContext(), id)
	if err != nil {
		return err
	}

	// incase of no error return response with 200 status
	return c.Status(http.StatusOK).JSON(resp)
}
