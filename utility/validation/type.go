package validation

import (
	"strconv"

	"github.com/go-playground/validator/v10"
)

func ValidatePositiveInteger(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	value, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return value > 0
}

func ValidateNonNegativeInteger(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	value, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return value >= 0
}

func ValidateBoolean(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	_, err := strconv.ParseBool(str)
	return err == nil
}
