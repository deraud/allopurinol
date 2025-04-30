package validation

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func ValidatePostalCode(fl validator.FieldLevel) bool {
	postalCode := fl.Field().String()

	if len(postalCode) != 5 {
		return false
	}
	for _, c := range postalCode {
		if !unicode.IsNumber(c) {
			return false
		}
	}
	return true
}
