package middlewares

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

var validate = validator.New()

func Validate[T any](target *T) fiber.Handler {
	return func(c *fiber.Ctx) error {
		newTarget := new(T)

		// 구조체의 필드 태그를 검사하여 파싱 방법 결정
		t := reflect.TypeOf(newTarget).Elem()
		hasQuery := false
		hasBody := false

		for i := 0; i < t.NumField(); i++ {
			if _, ok := t.Field(i).Tag.Lookup("query"); ok {
				hasQuery = true
			}
			if _, ok := t.Field(i).Tag.Lookup("json"); ok {
				hasBody = true
			}
		}

		// Query 파라미터 파싱
		if hasQuery {
			if err := c.QueryParser(newTarget); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Invalid query parameters",
				})
			}
		}

		// Request body 파싱
		if hasBody {
			if err := c.BodyParser(newTarget); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Invalid request body",
				})
			}
		}

		// Validation 실행
		if err := validate.Struct(newTarget); err != nil {
			// Validation 에러 처리
			var errors []ValidationError
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, ValidationError{
					Field: err.Field(),
					Tag:   err.Tag(),
					Value: err.Param(),
				})
			}
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"errors": errors,
			})
		}

		// Context에 검증된 데이터를 저장
		c.Locals("validatedData", newTarget)

		return c.Next()
	}
}
