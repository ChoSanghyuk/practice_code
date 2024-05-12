package basic

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func BasicRoute(app *fiber.App) {

	// GET http://localhost:3000/basic/hello%20world
	app.Get("/basic/:value", func(c fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
		// => Get request with value: hello world
	})

	// GET http://localhost:3000/baisc/john
	app.Get("/basic/:name?", func(c fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
			// => Hello john
		}
		return c.SendString("Where is john?")
	})

	// http://localhost:3000/basic/plantae/prunus.persica
	app.Get("/basic/plantae/:genus.:species", func(c fiber.Ctx) error {
		fmt.Fprintf(c, "%s.%s\n", c.Params("genus"), c.Params("species"))
		return nil // prunus.persica
	})

	// http://localhost:3000/basic/flights/LAX-SFO
	app.Get("/basic/flights/:from-:to", func(c fiber.Ctx) error {
		fmt.Fprintf(c, "%s-%s\n", c.Params("from"), c.Params("to"))
		return nil // LAX-SFO
	})

}
