package handler

import (
	"healthcare-allopurinol/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandlerImpl) LoginHandler(ctx *gin.Context) {
	var body dto.AuthLoginRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	token, err := h.service.LoginService(ctx, body.Email, body.Password)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: gin.H{
			"access_token": token,
		},
	})
}
