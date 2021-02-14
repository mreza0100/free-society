package connections

import (
	"fmt"
	"microServiceBoilerplate/configs"
	pb "microServiceBoilerplate/proto/generated/feed"

	"github.com/mreza0100/golog"
	"google.golang.org/grpc"
)

func FeedSrvConn(lgr *golog.Core) pb.FeedServiceClient {
	conn, err := grpc.Dial(configs.FeedConfigs.Addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(configs.PostConfigs.Timeout))
	if err != nil {
		fmt.Println("Cant connect to feed service")
		panic(err)
	}

	lgr.GreenLog("✅ Connected to feed service :)")
	return pb.NewFeedServiceClient(conn)
}
