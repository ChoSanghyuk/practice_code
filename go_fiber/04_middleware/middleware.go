package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func Logger(c *fiber.Ctx) error {

	err := c.Next()

	method := c.Method()
	path := c.Path()
	ip := c.IP()
	resq := string(c.Body())
	resp := string(c.Response().Body())

	log.Info(fmt.Sprintf("path : %s , method : %s ,  ip : %s , req : %s , resp : %s", path, method, ip, resq, resp))
	if err != nil {
		log.Error(fmt.Sprintf("path : %s , error :%s", path, err))
	}

	return nil
}

func Recover(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			log.Error(fmt.Sprintf("Server shutdown from %s", r))
		}
	}()
	return c.Next()
}

func Timeout(ts int) func(c *fiber.Ctx) error {

	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(ts)*time.Second)
		defer cancel()

		// 핸들러 고루틴에서 실행
		done := make(chan error, 1)
		go func() {
			done <- c.Next()
		}()

		// 핸들러의 응답 혹은 타임아웃을 기다림
		select {
		case err := <-done:
			// 핸들러가 끝났다면, 그대로 종료
			return err
		case <-ctx.Done():
			// 타임아웃 발생 시, ctx에 타임아웃 에러 저장하고 종료
			return c.Status(fiber.StatusRequestTimeout).SendString("Request timed out")
		}
	}
}
