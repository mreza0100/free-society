package connections

import (
	"fmt"
	"microServiceBoilerplate/configs"
	pb "microServiceBoilerplate/proto/generated/user"
	"time"

	"google.golang.org/grpc"
)

func UserSrvConn() pb.UserServiceClient {
	conn, err := grpc.Dial(configs.UserConfigs.Addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		fmt.Println("Cant connect to user service")
		panic(err)
	}
	fmt.Println("Connected to user service :)")
	return pb.NewUserServiceClient(conn)
}
