package group_route

import (
	"github.com/gofiber/fiber/v2"
)

func UseGroup(app *fiber.App) {

	api := app.Group("/group", middleware) // /group 시작되는 모든 url에 요청이 올 시, 실행되는 미들웨어 지정

	v1 := api.Group("/v1")   // /group/v1
	v1.Get("/list", handler) // /group/v1/list => URL PATH : /group/v1/list
	v1.Get("/user", handler) // /group/v1/user => URL PATH : /group/v1/user

	v2 := api.Group("/v2")   // /group/v2
	v2.Get("/list", handler) // /group/v2/list => URL PATH : /group/v2/list
	v2.Get("/user", handler) // /group/v2/user => URL PATH : /group/v2/user
}

func UseRoute(app *fiber.App) {
	app.Route("/route", func(api fiber.Router) {
		api.Get("/foo", handler).Name("foo") // /route/foo (name: test.foo) => URL PATH : /route/foo
		api.Get("/bar", handler).Name("bar") // /route/bar (name: test.bar) => URL PATH : /route/bar
	}, "test.")
}

func UseMount(app *fiber.App) {

	micro := fiber.New()
	app.Mount("/mount", micro)

	micro.Get("/doe", handler) // GET /mount/doe => URL PATH : /mount/doe

}
