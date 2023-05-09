package validation

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func registerCustomErrors() {
	customError["required"] = func(field validator.FieldError, translatedFieldName string) string {
		return fmt.Sprintf("'%s' tidak boleh kosong", translatedFieldName)
	}

	customError["uuid"] = func(field validator.FieldError, translatedFieldName string) string {
		return fmt.Sprintf("'%s' harus dengan format UUID", translatedFieldName)
	}

	customError["max"] = func(field validator.FieldError, translatedFieldName string) string {
		fieldType := field.Kind()
		if fieldType == reflect.Int {
			return fmt.Sprintf("'%s' tidak boleh lebih dari %s", translatedFieldName, field.Param())
		}
		return fmt.Sprintf("'%s' tidak boleh melebihi %s karakter", translatedFieldName, field.Param())
	}

	customError["numeric"] = func(field validator.FieldError, translatedFieldName string) string {
		return fmt.Sprintf("'%s' harus numeric", translatedFieldName)
	}

	customError["datetime"] = func(field validator.FieldError, translatedFieldName string) string {
		param := field.Param()
		if param == "2006-01-02" {
			param = "YYYY-MM-DD"
		}

		return fmt.Sprintf("'%s' harus dengan format %s", translatedFieldName, param)
	}

	customError["excludes"] = func(field validator.FieldError, translatedFieldName string) string {
		param := field.Param()
		if param == " " {
			param = "spasi"
		}

		return fmt.Sprintf("'%s' tidak boleh memiliki %s", translatedFieldName, param)
	}
}
