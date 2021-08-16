package connections

import (
	pb "freeSociety/proto/generated/relation"

	"freeSociety/services/relation/configs"

	"github.com/mreza0100/golog"
	"google.golang.org/grpc"
)

func RelationSrvConn(lgr *golog.Core) pb.RelationServiceClient {
	conn, err := grpc.Dial(configs.Configs.Addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(configs.Configs.ConnectionTimeout))
	if err != nil {
		lgr.Fatal("Cant connect to user service", err)
	}
	lgr.SuccessLog("Connected to relation service")
	return pb.NewRelationServiceClient(conn)
}
