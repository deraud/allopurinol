package utility

import "github.com/go-playground/validator/v10"

func ExtractErrorMessage(fe validator.FieldError) string {
	if fe.Tag() == "required" {
		return "this field is required"
	}
	if fe.Tag() == "email" {
		return "this field must be a valid email address"
	}
	if fe.Tag() == "max" {
		return "this field exceeds the maximum limit of " + fe.Param()
	}
	if fe.Tag() == "min" {
		return "this field's minimum value is " + fe.Param()
	}
	if fe.Tag() == "gte" {
		return "this field must be greater than or equal " + fe.Param()
	}
	if fe.Tag() == "gt" {
		return "this field must be greater than " + fe.Param()
	}
	if fe.Tag() == "password" {
		return "password must be between 8-50 characters and have at least one uppercase letter, number and special character"
	}
	if fe.Tag() == "positive" {
		return "this field must be a positive integer"
	}
	if fe.Tag() == "oneof" {
		return "this field must be one of the following values: " + fe.Param()
	}
	if fe.Tag() == "required_without" {
		return "one of this field is required"
	}
	if fe.Tag() == "e164" {
		return "this field must be a valid phone number in E.164 format"
	}
	if fe.Tag() == "non_negative" {
		return "this field must be a non negative integer"
	}
	if fe.Tag() == "assigned" {
		return "this field must be a boolean value"
	}
	if fe.Tag() == "positive_decimal" {
		return "this field must be a positive decimal"
	}
	if fe.Tag() == "datetime" {
		return "this field must be a valid time format"
	}
	if fe.Tag() == "required_if" {
		return "this field is required for the following value: " + fe.Param()
	}
	if fe.Tag() == "gtcsfield" {
		return "this field must be greater than: " + fe.Param()
	}
	if fe.Tag() == "size" {
		return "this field's size exceeds the maximum limit of " + fe.Param()
	}
	if fe.Tag() == "extension" {
		return "this field must be one of the following format: " + fe.Param()
	}
	return "unknown error"
}
