package connections

import (
	"microServiceBoilerplate/configs"
	pb "microServiceBoilerplate/proto/generated/feed"

	"github.com/mreza0100/golog"
	"google.golang.org/grpc"
)

func FeedSrvConn(lgr *golog.Core) pb.FeedServiceClient {
	conn, err := grpc.Dial(configs.FeedConfigs.Addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(configs.PostConfigs.Timeout))
	if err != nil {
		lgr.Fatal("Cant connect to feed service", err)
	}

	lgr.SuccessLog("Connected to feed service")
	return pb.NewFeedServiceClient(conn)
}
