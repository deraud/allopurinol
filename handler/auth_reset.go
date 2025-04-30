package handler

import (
	"healthcare-allopurinol/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandlerImpl) ForgotPasswordHandler(ctx *gin.Context) {
	var body dto.AuthForgotPasswordRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.service.ForgotPasswordService(ctx, body.Email)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *AuthHandlerImpl) ResetPasswordHandler(ctx *gin.Context) {
	var body dto.AuthResetPasswordRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.service.ResetPasswordService(ctx, body.Token, body.NewPassword)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
