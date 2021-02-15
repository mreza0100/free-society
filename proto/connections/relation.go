package connections

import (
	"microServiceBoilerplate/configs"
	pb "microServiceBoilerplate/proto/generated/relation"

	"github.com/mreza0100/golog"
	"google.golang.org/grpc"
)

func RelationSrvConn(lgr *golog.Core) pb.RelationServiceClient {
	conn, err := grpc.Dial(configs.RelationConfigs.Addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(configs.RelationConfigs.Timeout))
	if err != nil {
		lgr.Fatal("Cant connect to user service", err)
	}
	lgr.SuccessLog("Connected to relation service :)")
	return pb.NewRelationServiceClient(conn)
}
