package requestbody

import "github.com/gofiber/fiber/v2"

func RequestBodyRoute(app *fiber.App) {

	api := app.Group("/request_body")

	api.Get("/print", BodyHandler)
	api.Get("/parse", BodyParserHandler)
	api.Get("/valid_check", ValidateHandler)

}
