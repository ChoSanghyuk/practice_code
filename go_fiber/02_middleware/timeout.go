package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

var LimiterHandler func(*fiber.Ctx) error = limiter.New(limiter.Config{
	Max:        1, // number of request can be processed at a time
	Expiration: 2 * time.Second,
	LimitReached: func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusRequestTimeout).SendString("Request timed out")
	},
})
