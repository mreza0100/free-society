package connections

import (
	"microServiceBoilerplate/configs"
	pb "microServiceBoilerplate/proto/generated/post"

	"github.com/mreza0100/golog"
	"google.golang.org/grpc"
)

func PostSrvConn(lgr *golog.Core) pb.PostServiceClient {
	conn, err := grpc.Dial(configs.PostConfigs.Addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(configs.PostConfigs.Timeout))
	if err != nil {
		lgr.Fatal("Cant connect to post service", err)
	}

	lgr.SuccessLog("Connected to post service :)")
	return pb.NewPostServiceClient(conn)
}
