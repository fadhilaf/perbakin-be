package validation

import (
	"fmt"
	"reflect"
	"strings"

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

	customError["number"] = func(field validator.FieldError, translatedFieldName string) string {
		return fmt.Sprintf("'%s' harus angka", translatedFieldName)
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

	customError["boolean"] = func(field validator.FieldError, translatedFieldName string) string {
		return fmt.Sprintf("'%s' harus boolean", translatedFieldName)
	}

	customError["gte"] = func(field validator.FieldError, translatedFieldName string) string {
		return fmt.Sprintf("'%s' harus lebih dari %s", translatedFieldName, field.Param())
	}

	customError["lte"] = func(field validator.FieldError, translatedFieldName string) string {
		return fmt.Sprintf("'%s' harus kurang dari %s", translatedFieldName, field.Param())
	}

	customError["oneof"] = func(field validator.FieldError, translatedFieldName string) string {
		words := strings.Split(field.Param(), " ")
		lastWord := words[len(words)-1]

		param := strings.Join(words[:len(words)-1], ", ")
		param += ", atau " + lastWord

		return fmt.Sprintf("%s yang dipilih harus salah satu dari %s", translatedFieldName, param)
	}

	customError["len"] = func(field validator.FieldError, translatedFieldName string) string {
		return fmt.Sprintf("banyak item pada '%s' harus %s", translatedFieldName, field.Param())
	}
}
