package route

import (
	basic "go_fiber/01_basic"
	group_route "go_fiber/02_group_route"
	requestbody "go_fiber/03_request_body"
	middleware "go_fiber/04_middleware"

	"github.com/gofiber/fiber/v2"
)

func AddRoute(app *fiber.App) {

	app.Use(middleware.Logger)
	app.Use(middleware.Recover)
	app.Use(middleware.Timeout(3))

	// 01_basic
	basic.BasicRoute(app)

	// 02_group_route
	group_route.UseGroup(app)
	group_route.UseRoute(app)
	group_route.UseMount(app)

	// 03_request_body
	requestbody.RequestBodyRoute(app)
}
