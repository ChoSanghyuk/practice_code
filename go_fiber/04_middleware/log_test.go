package middleware

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

var logApp *fiber.App

func init() {

	logApp = fiber.New()

	logApp.Use(requestid.New())

	logApp.Use(logger.New(logger.Config{
		Format: "${pid} ${respHeader:X-Request-Id} ${status} - ${method} ${path} error: ${error}\n",
	}))

	logApp.Get("/log", func(c *fiber.Ctx) error {
		requestid := c.Locals("requestid").(string)
		// c.Request().Header
		fmt.Printf("Request Header %v\n", c.GetReqHeaders())
		fmt.Printf("Response Header %v\n", c.GetRespHeaders())
		return errors.New(requestid + "Error Occurred")
		// return c.SendString("Response form Log API. Request Id : " + requestid)
	})
}

func Test1(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/log", nil)
	resp, err := logApp.Test(req, -1)
	if err != nil {
		t.Error("Error Occurred", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Error("Response status should be 200")
	}

	body, _ := io.ReadAll(resp.Body)
	t.Log(string(body))
}
