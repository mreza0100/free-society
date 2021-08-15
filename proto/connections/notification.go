package connections

import (
	"freeSociety/configs"
	pb "freeSociety/proto/generated/notification"

	"github.com/mreza0100/golog"
	"google.golang.org/grpc"
)

func NotificationSrvConn(lgr *golog.Core) pb.NotificationServiceClient {
	conn, err := grpc.Dial(configs.NotificationConfigs.Addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(configs.NotificationConfigs.Timeout))
	if err != nil {
		lgr.Fatal("Cant connect to notification service", err)
	}

	lgr.SuccessLog("Connected to notification service")
	return pb.NewNotificationServiceClient(conn)
}
