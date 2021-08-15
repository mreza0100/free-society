package instances

import (
	pb "freeSociety/proto/generated/notification"
)

type Handlers interface {
	pb.NotificationServiceServer
}
