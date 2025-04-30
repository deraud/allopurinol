package validation

import (
	"healthcare-allopurinol/constant"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)

func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	var floor, ceiling, number, upper, special bool
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case strings.ContainsRune("\"!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~”", c):
			special = true
		}
		if c == ' ' {
			return false
		}
	}
	floor = len(password) >= constant.MIN_PASSWORD_LENGTH
	ceiling = len(password) <= constant.MAX_PASSWORD_LENGTH
	return floor && ceiling && number && upper && special
}
