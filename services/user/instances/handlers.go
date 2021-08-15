package instances

import (
	pb "freeSociety/proto/generated/user"
)

type Handlers interface {
	pb.UserServiceServer
}
