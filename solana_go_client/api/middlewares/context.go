package middlewares

import (
	"workspace/api/types"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func requestID(c *fiber.Ctx) error {
	requestID := c.Get("X-Request-ID")
	if requestID == "" {
		requestID = generateReqID()
		c.Set("X-Request-ID", requestID)
	}
	c.Locals(types.RequestID, requestID)

	return c.Next()
}

func generateReqID() string {
	return uuid.New().String()
}
