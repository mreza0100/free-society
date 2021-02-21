package instances

import (
	pb "microServiceBoilerplate/proto/generated/user"
	"microServiceBoilerplate/services/user/models"
)

type Sevice interface {
	NewUser(in *pb.NewUserRequest) (uint64, error)
	GetUser(id uint64, email string) (*pb.GetUserResponse, error)
	DeleteUser(id uint64) error
	GetUsers(ids []uint64) (map[uint64]*models.User, error)

	IsUserExist(userId uint64) bool
}
