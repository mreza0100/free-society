package main

import (
	"log"
	"microServiceBoilerplate/configs"
	pb "microServiceBoilerplate/proto/generated/user"
	"net"

	"google.golang.org/grpc"

	"github.com/mreza0100/golog"

	"microServiceBoilerplate/services/user/microservice"
)

func initLogger() *golog.Core {
	return golog.New(golog.InitOprions{
		LogPath:    "./services/user/logs/out.log",
		Name:       "User_Service",
		PanicOnErr: false,
	})
}

func main() {
	Lg := initLogger()
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, microservice.NewUserService(Lg))

	lis, err := net.Listen("tcp", ":"+configs.UserConfigs.StrPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
