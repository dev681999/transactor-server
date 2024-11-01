package api

import (
	"transactor-server/pkg/pkgerr"

	"github.com/gofiber/fiber/v2"
)

// ErrorHandler is a custom error handler which handles a pkgerr.HttpError if provided
// else use the default error handler
var ErrorHandler = func(c *fiber.Ctx, err error) error {
	if e, ok := err.(pkgerr.HttpError); ok {
		return c.Status(e.HttpStatusCode()).JSON(e.ResponseBody())
	}

	return fiber.DefaultErrorHandler(c, err)
}
