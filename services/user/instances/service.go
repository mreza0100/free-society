package instances

import (
	pb "freeSociety/proto/generated/user"
	"freeSociety/services/user/models"
)

type Sevice interface {
	NewUser(name, email, gender, avatarFormat string, avatar []byte) (uint64, error)
	GetUser(requestor, id uint64, email string) (*pb.GetUserResponse, error)
	DeleteUser(id uint64) error
	GetUsers(ids []uint64) (map[uint64]*models.User, error)
	UpdateUser(userId uint64, name, gender, avatarFormat string, avatar []byte) error

	IsUserExist(userId uint64) bool
}
