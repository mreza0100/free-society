package types

import (
	pb "microServiceBoilerplate/proto/generated/user"
)

type Handlers interface {
	pb.UserServiceServer

	IsUserExist(userId uint64) bool
}
