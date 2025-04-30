package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/sentinel"
	"healthcare-allopurinol/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandlerImpl) VerifyHandler(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	err := h.service.VerifyService(ctx, content.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *AuthHandlerImpl) VerifyTokenHandler(ctx *gin.Context) {
	var body dto.AuthVerifyTokenRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.service.VerifyTokenService(ctx, body.Token)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
