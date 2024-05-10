package main

import "github.com/gofiber/fiber/v3"

func main() {
	app := fiber.New()

	app.Get("/:name.:age", func(c fiber.Ctx) error {
		return c.SendString("Hello!" + c.Params("name") + "   " + c.Params("age")) //+
	})

	app.Listen(":3000")
}
