package main

import (
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
	return golog.New(golog.InitOprions{
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

	lis, err := net.Listen("tcp", ":"+configs.UserConfigs.StrPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
