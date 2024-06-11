package route

import (
	basic "go_fiber/01_basic"
	group_route "go_fiber/02_group_route"
	requestbody "go_fiber/03_request_body"

	"github.com/gofiber/fiber/v2"
)

func AddRoute(app *fiber.App) {

	// 01_basic
	basic.BasicRoute(app)

	// 02_group_route
	group_route.UseGroup(app)
	group_route.UseRoute(app)
	group_route.UseMount(app)

	// 03_request_body
	requestbody.RequestBodyRoute(app)
}
