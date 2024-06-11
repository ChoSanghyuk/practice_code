package requestbody

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ValidCheck(c *fiber.Ctx, s any) error {
	if errs := myValidator.Validate(s); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)
		for _, err := range errs {
			if err.Tag == "required" {
				errMsgs = append(errMsgs, fmt.Sprintf(
					"필수 필드 %s 누락",
					err.FailedField,
				))
			} else {
				errMsgs = append(errMsgs, fmt.Sprintf(
					"%s 타입의 필드 %s가 올바르지 않음",
					err.Tag,
					err.FailedField,
				))
			}
		}
		return errors.New(strings.Join(errMsgs, ", "))
	}
	return nil
}
