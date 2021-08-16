package connections

import (
	pb "freeSociety/proto/generated/security"

	"freeSociety/services/security/configs"

	"github.com/mreza0100/golog"
	"google.golang.org/grpc"
)

func SecuritySrvConn(lgr *golog.Core) pb.SecurityServiceClient {
	conn, err := grpc.Dial(configs.Configs.Addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(configs.Configs.ConnectionTimeout))
	if err != nil {
		lgr.Fatal("Cant connect to security service", err)
	}
	lgr.SuccessLog("Connected to security service")
	return pb.NewSecurityServiceClient(conn)
}
