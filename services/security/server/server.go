package main

import (
	"log"
	"microServiceBoilerplate/configs"
	pb "microServiceBoilerplate/proto/generated/security"
	"microServiceBoilerplate/utils"
	"net"

	"google.golang.org/grpc"

	"github.com/mreza0100/golog"

	"microServiceBoilerplate/services/security/microservice"
)

func initLogger() *golog.Core {
	return golog.New(golog.InitOpns{
		LogPath:   configs.LogPath,
		Name:      "Security_Service",
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
	pb.RegisterSecurityServiceServer(grpcServer, service)

	lis, err := net.Listen("tcp", ":"+configs.UserConfigs.StrPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
