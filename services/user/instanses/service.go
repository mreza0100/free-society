package instanses

import (
	pb "microServiceBoilerplate/proto/generated/user"
)

type Sevice interface {
	NewUser(in *pb.NewUserRequest) (uint64, error)
	GetUser(id uint64, email string) (*pb.GetUserResponse, error)
	DeleteUser(id uint64) error

	IsUserExist(userId uint64) bool
}
