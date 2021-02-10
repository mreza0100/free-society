package connections

import (
	"fmt"
	"microServiceBoilerplate/configs"
	pb "microServiceBoilerplate/proto/generated/post"

	"google.golang.org/grpc"
)

func PostSrvConn() pb.PostServiceClient {
	conn, err := grpc.Dial(configs.PostConfigs.Addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(configs.PostConfigs.Timeout))
	if err != nil {
		fmt.Println("Cant connect to post service")
		panic(err)
	}

	fmt.Println("✅ Connected to post service :)")
	return pb.NewPostServiceClient(conn)
}
