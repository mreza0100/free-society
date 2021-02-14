package main

import (
	"log"
	"microServiceBoilerplate/configs"
	pb "microServiceBoilerplate/proto/generated/feed"
	"microServiceBoilerplate/utils"
	"net"

	"google.golang.org/grpc"

	"github.com/mreza0100/golog"

	"microServiceBoilerplate/services/feed/microservice"
)

func initLogger() *golog.Core {
	return golog.New(golog.InitOpns{
		LogPath:   configs.LogPath,
		Name:      "Feed_Service",
		WithTime:  true,
		DebugMode: utils.IsDevMode,
	})
}

func main() {
	var (
		lgr        = initLogger()
		service    = microservice.NewFeedService(lgr)
		grpcServer = grpc.NewServer()
	)
	pb.RegisterFeedServiceServer(grpcServer, service)

	lis, err := net.Listen("tcp", ":"+configs.FeedConfigs.StrPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	lgr.GreenLog("port is open now")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
