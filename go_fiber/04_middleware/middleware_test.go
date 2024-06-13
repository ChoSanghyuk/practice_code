package middleware

import (
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var LimitTestApp *fiber.App

func init() {

	LimitTestApp = fiber.New()

	// LimitTestApp.Use(LimiterHandler)

	timeoutApi := LimitTestApp.Group("timeout", Timeout)

	timeoutApi.Get("/do", func(c *fiber.Ctx) error {
		log.Info("timeout called")
		time.Sleep(5 * time.Second)
		return c.SendString("Timeout didn't work")
	})

	timeoutApi.Get("/no", func(c *fiber.Ctx) error {
		log.Info("no_timeout called")
		return c.SendString("Timeout didn't work")
	})

	recoverApi := LimitTestApp.Group("recover", Recover)
	recoverApi.Get("/do", func(c *fiber.Ctx) error {
		log.Info("recover called")
		panic("Panic")
		return c.SendString("Panic didn't work")
	})

	loggerApi := LimitTestApp.Group("logging", logger.New())
	loggerApi.Get("/do", func(c *fiber.Ctx) error {
		log.Info("log called")
		return c.SendString("Response form Log API")
	})
}

func TestTimeout(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/timeout/do", nil)
	resp, err := LimitTestApp.Test(req, -1)
	if err != nil {
		t.Error("Error Occurred", err)
	}
	if resp.StatusCode != fiber.StatusRequestTimeout {
		t.Error("Response status should be 408")
	}

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "Request timed out" {
		t.Error("Response body should be 'Request timed out'")
	}

	t.Log(string(body))
}

func TestTimeoutNotWork(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/timeout/no", nil)
	resp, err := LimitTestApp.Test(req, -1)
	if err != nil {
		t.Error("Error Occurred", err)
	}

	body, _ := io.ReadAll(resp.Body)
	if string(body) == "Request timed out" {
		t.Error("Response body should not be 'Request timed out'")
	}

	t.Log(string(body))
}

func TestRecover(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/recover/do", nil)
	resp, err := LimitTestApp.Test(req, -1)
	if err != nil {
		t.Error("Error Occurred", err)
	}

	body, _ := io.ReadAll(resp.Body)
	t.Log(string(body))
}

func TestLogger(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/logging/do", strings.NewReader("log API Request"))
	resp, err := LimitTestApp.Test(req, -1)
	if err != nil {
		t.Error("Error Occurred", err)
	}

	body, _ := io.ReadAll(resp.Body)
	t.Log(string(body))
}
