package route

import (
	basic "go_fiber/01_basic"

	"github.com/gofiber/fiber/v3"
)

func AddRoute(app *fiber.App) {

	basic.BasicRoute(app)
}