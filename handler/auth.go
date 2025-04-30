package handler

import (
	"healthcare-allopurinol/service"

	"github.com/gin-gonic/gin"
)

type AuthHandlerItf interface {
	LoginHandler(ctx *gin.Context)
	RegisterHandler(ctx *gin.Context)
	ForgotPasswordHandler(ctx *gin.Context)
	ResetPasswordHandler(ctx *gin.Context)
	VerifyHandler(ctx *gin.Context)
	VerifyTokenHandler(ctx *gin.Context)
}

type AuthHandlerImpl struct {
	service service.AuthServiceItf
}

func NewAuthHandler(service service.AuthServiceItf) AuthHandlerItf {
	return &AuthHandlerImpl{service: service}
}
