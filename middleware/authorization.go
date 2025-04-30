package middleware

import (
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/sentinel"
	"healthcare-allopurinol/utility"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrUnauthorized)
		ctx.Abort()
		return
	}
	content := c.(*utility.ClaimsContent)

	if content.Role != constant.ROLE_ADMIN {
		ctx.Error(sentinel.ErrUnauthorized)
		ctx.Abort()
		return
	}

	ctx.Next()
}

func UserMiddleware(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrUnauthorized)
		ctx.Abort()
		return
	}
	content := c.(*utility.ClaimsContent)

	if content.Role != constant.ROLE_USER {
		ctx.Error(sentinel.ErrUnauthorized)
		ctx.Abort()
		return
	}

	ctx.Next()
}

func PharmacistMiddleware(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrUnauthorized)
		ctx.Abort()
		return
	}
	content := c.(*utility.ClaimsContent)

	if content.Role != constant.ROLE_PHARMACIST {
		ctx.Error(sentinel.ErrUnauthorized)
		ctx.Abort()
		return
	}

	ctx.Next()
}
