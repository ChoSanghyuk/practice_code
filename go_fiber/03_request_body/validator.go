package requestbody

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       any
}

var myValidator = validator.New()

func init() {
	myValidator.RegisterValidation("teener", func(fl validator.FieldLevel) bool {
		return fl.Field().Int() >= 12 && fl.Field().Int() <= 18
	})
}

func validCheck(s any) error {
	if errs := validate(s); len(errs) > 0 && errs[0].Error {
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

func validate(data any) []ErrorResponse {
	validationErrors := []ErrorResponse{}
	errs := myValidator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}
	return validationErrors
}
