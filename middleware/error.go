package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/sentinel"
	"healthcare-allopurinol/utility"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorMiddleware(ctx *gin.Context) {
	ctx.Next()

	if len(ctx.Errors) == 0 {
		return
	}

	err := ctx.Errors[0]

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		handleValidationError(ctx, ve)
		return
	}

	var ute *json.UnmarshalTypeError
	if errors.As(err, &ute) {
		handleUnmarshalTypeError(ctx, ute)
		return
	}

	var ue sentinel.UnauthorizedError
	if errors.As(err, &ue) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Error: []sentinel.ErrorWrapper{ue.ErrorWrapper},
		})
		return
	}

	var fe sentinel.ForbiddenError
	if errors.As(err, &fe) {
		ctx.AbortWithStatusJSON(http.StatusForbidden, dto.Response{
			Error: []sentinel.ErrorWrapper{fe.ErrorWrapper},
		})
		return
	}

	var bre sentinel.BadRequestError
	if errors.As(err, &bre) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			Error: []sentinel.ErrorWrapper{bre.ErrorWrapper},
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
		Error: []sentinel.ErrorWrapper{
			{
				Field:   "server",
				Message: "internal server error",
			},
		},
	})
}

func handleValidationError(ctx *gin.Context, ve validator.ValidationErrors) {
	var fieldErrors []sentinel.ErrorWrapper
	for _, fe := range ve {
		fieldErrors = append(fieldErrors, sentinel.ErrorWrapper{
			Field:   strings.ToLower(fe.Field()),
			Message: utility.ExtractErrorMessage(fe),
		})
	}
	ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{Error: fieldErrors})
}

func handleUnmarshalTypeError(ctx *gin.Context, ute *json.UnmarshalTypeError) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
		Error: []sentinel.ErrorWrapper{
			{
				Field:   ute.Field,
				Message: fmt.Sprintf("this field should not be %s", ute.Value),
			},
		},
	})
}
