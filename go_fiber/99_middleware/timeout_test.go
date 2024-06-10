package middleware

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
)

var LimitTestApp *fiber.App

func init() {

	LimitTestApp = fiber.New()

	// LimitTestApp.Use(LimiterHandler)

	LimitTestApp.Use(Timeout)

	LimitTestApp.Get("/timeout", func(c *fiber.Ctx) error {
		fmt.Println("/timeout called")
		time.Sleep(5 * time.Second)
		return c.SendString("Timeout didn't work")
	})

	LimitTestApp.Get("/no_timeout", func(c *fiber.Ctx) error {
		fmt.Println("/no_timeout called")
		return c.SendString("Timeout didn't work")
	})
}

func TestTimeout(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/timeout", nil)
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
	req, _ := http.NewRequest(http.MethodGet, "/no_timeout", nil)
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
