package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandlerImpl) RegisterHandler(ctx *gin.Context) {
	var body dto.AuthRegisterRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	const USER = 2
	err = h.service.RegisterService(ctx, &entity.Credential{
		Email:    body.Email,
		Password: body.Password,
		RoleId:   USER,
	}, &entity.User{
		Name: body.Name,
	})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
