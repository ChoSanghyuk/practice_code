package middleware

import (
	"context"
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

	LimitTestApp.Use(func(c *fiber.Ctx) error {
		fmt.Println("middleware called")
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		// time.Sleep(2 * time.Second)
		// Create a new Fiber context that can be cancelled
		// fc := c.Clone()

		// Run the handler in a goroutine so we can use select
		done := make(chan bool, 1)
		go func() {
			c.Next() // Run the next handler
			done <- true
		}()

		// Wait for the handler to finish or the context to be cancelled
		select {
		case <-done:
			// If the handler finished, just return nil
			return nil
		case <-ctx.Done():
			// If the context was cancelled, return a timeout error
			return c.Status(fiber.StatusRequestTimeout).SendString("Request timed out")
		}
	})

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
