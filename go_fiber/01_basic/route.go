package basic

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func BasicRoute(app *fiber.App) {

	// GET http://localhost:3000/basic/param/hello%20world
	app.Get("/basic/param/:value", func(c fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
		// => Get request with value: hello world
	})

	// GET http://localhost:3000/basic/param_option/john
	app.Get("/basic/param_option/:name?", func(c fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
			// => Hello john
		}
		return c.SendString("Where is john?")
	})

	// http://localhost:3000/basic/multi_param1/prunus.persica
	app.Get("/basic/multi_param1/:genus.:species", func(c fiber.Ctx) error {
		fmt.Fprintf(c, "%s.%s\n", c.Params("genus"), c.Params("species"))
		return nil // prunus.persica
	})

	// http://localhost:3000/basic/multi_param2/LAX-SFO
	app.Get("/basic/multi_param2/:from-:to", func(c fiber.Ctx) error {
		fmt.Fprintf(c, "%s-%s\n", c.Params("from"), c.Params("to"))
		return nil // LAX-SFO
	})

	// http://localhost:3000/basic/multi_param2/LAX-SFO
	app.Static("/basic/static", "./static")
}
