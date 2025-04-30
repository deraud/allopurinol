package middleware

import (
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/sentinel"
	"healthcare-allopurinol/utility"

	"github.com/gin-gonic/gin"
)

func VerificationMiddleware(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrUnauthorized)
		ctx.Abort()
		return
	}
	content := c.(*utility.ClaimsContent)

	if content.Role == constant.ROLE_USER && !content.IsVerified {
		ctx.Error(sentinel.ErrUnverified)
		ctx.Abort()
		return
	}

	ctx.Next()
}
