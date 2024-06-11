package group_route

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// write a handler function that responds with a url of itself

func middleware(c *fiber.Ctx) error {
	log.Println("I am a middleware")
	return c.Next()
}

func handler(c *fiber.Ctx) error {
	return c.SendString("URL PATH : " + c.OriginalURL())
}
