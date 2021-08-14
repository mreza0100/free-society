package connections

import (
	"microServiceBoilerplate/configs"
	pb "microServiceBoilerplate/proto/generated/security"

	"github.com/mreza0100/golog"
	"google.golang.org/grpc"
)

func SecuritySrvConn(lgr *golog.Core) pb.SecurityServiceClient {
	conn, err := grpc.Dial(configs.SecurityConfigs.Addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(configs.RelationConfigs.Timeout))
	if err != nil {
		lgr.Fatal("Cant connect to security service", err)
	}
	lgr.SuccessLog("Connected to security service")
	return pb.NewSecurityServiceClient(conn)
}
