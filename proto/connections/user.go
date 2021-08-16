package connections

import (
	pb "freeSociety/proto/generated/user"
	"freeSociety/services/notification/configs"

	"github.com/mreza0100/golog"
	"google.golang.org/grpc"
)

func UserSrvConn(lgr *golog.Core) pb.UserServiceClient {
	conn, err := grpc.Dial(configs.Configs.Addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(configs.Configs.ConnectionTimeout))
	if err != nil {
		lgr.Fatal("Cant connect to user service", err)
	}
	lgr.SuccessLog("Connected to user service")
	return pb.NewUserServiceClient(conn)
}
