package requestbody

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       any
}

var myValidator = validator.New()

func Validate(data any) []ErrorResponse {
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

func init() {
	myValidator.RegisterValidation("teener", func(fl validator.FieldLevel) bool {
		return fl.Field().Int() >= 12 && fl.Field().Int() <= 18
	})
}
