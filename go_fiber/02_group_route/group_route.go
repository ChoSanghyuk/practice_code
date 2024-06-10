package group_route

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func GroupRoute(app *fiber.App) {

	api := app.Group("/api", middleware) // /api로 시작되는 모든 url에 요청이 올 시, 실행되는 미들웨어 지정

	v1 := api.Group("/v1")   // /api/v1
	v1.Get("/list", handler) // /api/v1/list
	v1.Get("/user", handler) // /api/v1/user

	v2 := api.Group("/v2")   // /api/v2
	v2.Get("/list", handler) // /api/v2/list
	v2.Get("/user", handler) // /api/v2/user

	log.Fatal(app.Listen(":3000"))

}
