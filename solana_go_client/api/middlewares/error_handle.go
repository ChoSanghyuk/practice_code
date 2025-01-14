package middlewares

import (
	"workspace/api/parameters"

	"github.com/gofiber/fiber/v2"
)

func error_handle(c *fiber.Ctx) error {
	err := c.Next()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			parameters.NewErrorResponse(
				err.Error(),
			),
		)
	}
	return nil
}
