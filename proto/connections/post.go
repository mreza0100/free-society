package connections

import (
	"fmt"
	"microServiceBoilerplate/configs"
	pb "microServiceBoilerplate/proto/generated/post"

	"github.com/mreza0100/golog"
	"google.golang.org/grpc"
)

func PostSrvConn(lgr *golog.Core) pb.PostServiceClient {
	conn, err := grpc.Dial(configs.PostConfigs.Addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(configs.PostConfigs.Timeout))
	if err != nil {
		fmt.Println("Cant connect to post service")
		panic(err)
	}

	lgr.GreenLog("✅ Connected to post service :)")
	return pb.NewPostServiceClient(conn)
}
