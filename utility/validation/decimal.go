package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

func ValidatePositiveDecimal(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	value, err := decimal.NewFromString(str)
	if err != nil {
		return false
	}
	return value.GreaterThan(decimal.Zero)
}
