package main

import (
	"fmt"
	pb "freeSociety/proto/generated/notification"
	"freeSociety/utils"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/mreza0100/golog"

	"freeSociety/services/notification/configs"
	"freeSociety/services/notification/microservice"
)

func initLogger() *golog.Core {
	return golog.New(golog.InitOpns{
		LogPath:   configs.Configs.LogPath,
		Name:      "notification_Service",
		WithTime:  true,
		DebugMode: utils.IsDevMode,
	})
}

func main() {
	var (
		lgr        = initLogger()
		service    = microservice.NewNotificationService(lgr)
		grpcServer = grpc.NewServer()
	)
	pb.RegisterNotificationServiceServer(grpcServer, service)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", configs.Configs.Service_port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
