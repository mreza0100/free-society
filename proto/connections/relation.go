package connections

import (
	"fmt"
	"microServiceBoilerplate/configs"
	pb "microServiceBoilerplate/proto/generated/relation"

	"github.com/mreza0100/golog"
	"google.golang.org/grpc"
)

func RelationSrvConn(lgr *golog.Core) pb.RelationServiceClient {
	conn, err := grpc.Dial(configs.RelationConfigs.Addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(configs.RelationConfigs.Timeout))
	if err != nil {
		fmt.Println("Cant connect to user service")
		panic(err)
	}
	lgr.GreenLog("✅ Connected to relation service :)")
	return pb.NewRelationServiceClient(conn)
}
