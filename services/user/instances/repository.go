package instances

import "freeSociety/services/user/models"

type (
	read interface {
		GetUserById(id uint64) (*models.User, error)
		GetUserByEmail(email string) (*models.User, error)
		IsUserExistById(userId uint64) bool
		IsUserExistByEmail(email string) bool
		GetUsersByIds(userIds []uint64) ([]*models.User, error)
	}
	write interface {
		NewUser(name, gender, email, avatarPath string) (uint64, error)
		DeleteUser(userId uint64) (picturePath string, err error)
		UpdateUser(userId uint64, name, gender, avatarPath string) error
	}

	Repository struct {
		Read  read
		Write write
	}
)
