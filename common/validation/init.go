package validation

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var customError map[string]func(field validator.FieldError, translatedFieldName string) string

var customMessage map[string]string

func InitValidation(validate *validator.Validate) {
	customMessage = map[string]string{}
	registerCustomMessages()

	customError = map[string]func(validator.FieldError, string) string{}
	registerCustomErrors()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}
