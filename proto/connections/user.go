package connections

import (
	"microServiceBoilerplate/configs"
	pb "microServiceBoilerplate/proto/generated/user"

	"github.com/mreza0100/golog"
	"google.golang.org/grpc"
)

func UserSrvConn(lgr *golog.Core) pb.UserServiceClient {
	conn, err := grpc.Dial(configs.UserConfigs.Addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(configs.UserConfigs.Timeout))
	if err != nil {
		lgr.Fatal("Cant connect to user service", err)
	}
	lgr.SuccessLog("Connected to user service")
	return pb.NewUserServiceClient(conn)
}
