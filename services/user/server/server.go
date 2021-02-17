package main

import (
	"fmt"
	"log"
	"microServiceBoilerplate/configs"
	pb "microServiceBoilerplate/proto/generated/user"
	"microServiceBoilerplate/utils"
	"net"

	"google.golang.org/grpc"

	"github.com/mreza0100/golog"

	"microServiceBoilerplate/services/user/microservice"
)

func initLogger() *golog.Core {
	return golog.New(golog.InitOpns{
		LogPath:   configs.LogPath,
		Name:      "User_Service",
		WithTime:  true,
		DebugMode: utils.IsDevMode,
	})
}

func main() {
	var (
		lgr        = initLogger()
		service    = microservice.NewUserService(lgr)
		grpcServer = grpc.NewServer()
	)
	pb.RegisterUserServiceServer(grpcServer, service)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", configs.UserConfigs.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
