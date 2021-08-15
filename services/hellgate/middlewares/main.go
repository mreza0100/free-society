package middlewares

import (
	"freeSociety/services/hellgate/security"
	"freeSociety/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetWritter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var w http.ResponseWriter = ctx.Writer

		utils.SetValGinCtx(ctx, security.WRITE_KEY_CTX, &w)

		ctx.Next()
	}
}

func SetRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		utils.SetValGinCtx(ctx, security.REQUEST_KEY_CTX, ctx.Request)

		ctx.Next()
	}
}
