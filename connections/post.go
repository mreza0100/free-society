package connections

import (
	pb "freeSociety/proto/generated/post"
	"freeSociety/services/post/configs"

	"github.com/mreza0100/golog"
	"google.golang.org/grpc"
)

func PostSrvConn(lgr *golog.Core) pb.PostServiceClient {
	conn, err := grpc.Dial(configs.Configs.Addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(configs.Configs.ConnectionTimeout))
	if err != nil {
		lgr.Fatal("Cant connect to post service", err)
	}

	lgr.SuccessLog("Connected to post service")
	return pb.NewPostServiceClient(conn)
}
