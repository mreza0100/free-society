package main

import (
	"fmt"
	"microServiceBoilerplate/services/hellgate/graph"
	"microServiceBoilerplate/services/hellgate/graph/generated"

	"microServiceBoilerplate/proto/connections"
	"microServiceBoilerplate/proto/generated/user"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func getConnections() user.UserServiceClient {
	return connections.UserSrvConn()
}

func main() {
	graphqlServer := handler.NewDefaultServer(generated.NewExecutableSchema(graph.New(getConnections())))
	ginServer := gin.New()
	ginServer.Use(gin.Recovery())

	ginServer.GET("/", func(ctx *gin.Context) {
		playground.Handler("todo", "/")(ctx.Writer, ctx.Request)
	})
	ginServer.POST("/", func(ctx *gin.Context) {
		graphqlServer.ServeHTTP(ctx.Writer, ctx.Request)
	})
	fmt.Println("⭕Hellgate is open now⭕")
	ginServer.Run(":10000")
}
