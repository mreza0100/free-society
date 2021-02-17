package main

import (
	"fmt"
	"log"
	"microServiceBoilerplate/configs"
	pb "microServiceBoilerplate/proto/generated/post"
	"microServiceBoilerplate/utils"
	"net"

	"google.golang.org/grpc"

	"github.com/mreza0100/golog"

	"microServiceBoilerplate/services/post/microservice"
)

func initLogger() *golog.Core {
	return golog.New(golog.InitOpns{
		LogPath:   configs.LogPath,
		Name:      "Post_Service",
		WithTime:  true,
		DebugMode: utils.IsDevMode,
	})
}

func main() {
	var (
		Lgr        = initLogger()
		service    = microservice.NewPostService(Lgr)
		grpcServer = grpc.NewServer()
	)
	pb.RegisterPostServiceServer(grpcServer, service)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", configs.PostConfigs.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
