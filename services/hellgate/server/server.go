package main

import (
	"context"
	"microServiceBoilerplate/configs"
	"microServiceBoilerplate/services/hellgate/graph"
	"microServiceBoilerplate/services/hellgate/graph/generated"
	"microServiceBoilerplate/services/hellgate/security"
	"microServiceBoilerplate/utils"

	"microServiceBoilerplate/proto/connections"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/mreza0100/golog"
)

func initLogger() *golog.Core {
	return golog.New(golog.InitOpns{
		LogPath:   configs.LogPath,
		Name:      "🔥__Hellgate__🔥",
		WithTime:  true,
		DebugMode: utils.IsDevMode,
	})
}

func main() {
	lgr := initLogger()
	graphqlServer := handler.NewDefaultServer(generated.NewExecutableSchema(*graph.New(graph.NewOpts{
		Lgr: lgr,

		RelationConn: connections.RelationSrvConn(lgr),
		UserConn:     connections.UserSrvConn(lgr),
		PostConn:     connections.PostSrvConn(lgr),
	})))

	ginServer := gin.New()
	ginServer.Use(gin.Recovery())

	ginServer.GET("/", func(ctx *gin.Context) {
		playground.Handler("micro", "/")(ctx.Writer, ctx.Request)
	})

	ginServer.Use(security.Middleware())

	ginServer.Use(func(ctx *gin.Context) {
		set := ctx.SetCookie
		newCtx := context.WithValue(ctx.Request.Context(), "mamad", set)

		ctx.Request = ctx.Request.WithContext(newCtx)
		ctx.Next()
	})

	ginServer.POST("/", func(ctx *gin.Context) {
		graphqlServer.ServeHTTP(ctx.Writer, ctx.Request)
	})

	lgr.Log("🔥🔥🔥 Hellgate is open now 🔥🔥🔥")
	ginServer.Run(":" + configs.HellgateConfigs.StrPort)
}
