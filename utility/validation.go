package utility

import (
	"healthcare-allopurinol/utility/validation"
	"mime/multipart"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var validationFunctions = map[string]func(fl validator.FieldLevel) bool{
	"password":     validation.ValidatePassword,
	"positive":     validation.ValidatePositiveInteger,
	"non_negative": validation.ValidateNonNegativeInteger,
	"boolean":      validation.ValidateBoolean,
	"postal_code":  validation.ValidatePostalCode,
	"positive_decimal": validation.ValidatePositiveDecimal,
}

func InitValidator() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		err := registerValidation(v, validationFunctions)
		if err != nil {
			return err
		}

		v.RegisterStructValidation(validation.ValidateImage, multipart.FileHeader{})
	}
	return nil
}

func registerValidation(v *validator.Validate, functions map[string]func(validator.FieldLevel) bool) error {
	for name, function := range functions {
		err := v.RegisterValidation(name, function)
		if err != nil {
			return err
		}
	}
	return nil
}
