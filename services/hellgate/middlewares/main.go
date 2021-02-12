package middlewares

import (
	"microServiceBoilerplate/services/hellgate/security"
	"microServiceBoilerplate/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetWritter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var w http.ResponseWriter = ctx.Writer

		utils.SetValGinCtx(ctx, security.WriterKeyCtx, &w)

		ctx.Next()
	}
}

func SetRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		utils.SetValGinCtx(ctx, security.RequestKeyCtx, ctx.Request)

		ctx.Next()
	}
}
