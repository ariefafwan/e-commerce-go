package request

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(data interface{}) map[string]string {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		field := e.Field()
		switch e.Tag() {
		case "required":
			errors[field] = field + " wajib diisi"
		case "gte":
			errors[field] = field + " minimal " + e.Param()
		case "lte":
			errors[field] = field + " maksimal " + e.Param()
		default:
			errors[field] = "Field " + strings.ToLower(field) + " tidak valid"
		}
	}
	return errors
}
