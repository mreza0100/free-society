package main

import (
	"fmt"
	"freeSociety/configs"
	pb "freeSociety/proto/generated/security"
	"freeSociety/utils"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/mreza0100/golog"

	"freeSociety/services/security/microservice"
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
		service    = microservice.NewSecurityService(lgr)
		grpcServer = grpc.NewServer()
	)
	pb.RegisterSecurityServiceServer(grpcServer, service)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", configs.SecurityConfigs.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
