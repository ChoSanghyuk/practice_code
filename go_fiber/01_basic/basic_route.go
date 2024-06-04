package basic

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func BasicRoute(app *fiber.App) {

	// GET http://localhost:3000/basic/param/hello_world
	app.Get("/basic/param/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
		// => Get request with value: hello_world
	})

	// GET http://localhost:3000/basic/param_option/john
	app.Get("/basic/param_option/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
			// => Hello john
		}
		return c.SendString("Where is john?")
	})

	// http://localhost:3000/basic/multi_param1/prunus.persica
	app.Get("/basic/multi_param1/:genus.:species", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("genus is %s and persica is %s", c.Params("genus"), c.Params("species")))
		// => genus is prunus and persica is persica
	})

	// http://localhost:3000/basic/multi_param2/LAX-SFO
	app.Get("/basic/multi_param2/:from-:to", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("from is %s and to is %s", c.Params("from"), c.Params("to")))
		// => from is LAX and to is SFO
	})

	// http://localhost:3000/basic/static/sample.html
	app.Static("/basic/static", "./01_basic/static")

	app.Get("/basic/static/sample", func(c *fiber.Ctx) error {
		return c.SendFile("./01_basic/static/sample.html")
	})

}
