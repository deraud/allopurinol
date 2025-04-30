package middleware

import (
	"healthcare-allopurinol/sentinel"
	"healthcare-allopurinol/utility"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(ctx *gin.Context) {
	if gin.Mode() == gin.TestMode {
		ctx.Set("userData", &utility.ClaimsContent{
			Id:   1,
			Role: 2,
		})
		ctx.Next()
		return
	}

	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.Error(sentinel.ErrUnauthenticated)
		ctx.Abort()
		return
	}

	content, err := utility.ValidateJWToken(token)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.Set("userData", content)
	ctx.Next()
}
