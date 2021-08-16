package connections

import (
	pb "freeSociety/proto/generated/feed"
	"freeSociety/services/feed/configs"

	"github.com/mreza0100/golog"
	"google.golang.org/grpc"
)

func FeedSrvConn(lgr *golog.Core) pb.FeedServiceClient {
	conn, err := grpc.Dial(configs.Configs.Addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(configs.Configs.ConnectionTimeout))
	if err != nil {
		lgr.Fatal("Cant connect to feed service", err)
	}

	lgr.SuccessLog("Connected to feed service")
	return pb.NewFeedServiceClient(conn)
}
