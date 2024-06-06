package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

var ts = time.Duration(3)

func Timeout(c *fiber.Ctx) error {
	fmt.Println("middleware called")
	ctx, cancel := context.WithTimeout(context.Background(), ts*time.Second)
	defer cancel()

	// 핸들러 고루틴에서 실행
	done := make(chan bool, 1)
	go func() {
		c.Next()
		done <- true
	}()

	// 핸들러의 응답 혹은 타임아웃을 기다림
	select {
	case <-done:
		// 핸들러가 끝났다면, 그대로 종료
		return nil
	case <-ctx.Done():
		// 타임아웃 발생 시, ctx에 타임아웃 에러 저장하고 종료
		return c.Status(fiber.StatusRequestTimeout).SendString("Request timed out")
	}
}
