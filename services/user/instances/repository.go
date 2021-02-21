package instances

import "microServiceBoilerplate/services/user/models"

type (
	read interface {
		GetUserById(id uint64) (*models.User, error)
		GetUserByEmail(email string) (*models.User, error)
		IsUserExistById(userId uint64) bool
		IsUserExistByEmail(email string) bool
		GetUsersByIds(userIds []uint64) ([]*models.User, error)
	}
	write interface {
		NewUser(user *models.User) (uint64, error)
		DeleteUser(id uint64) error
	}

	Repository struct {
		Read  read
		Write write
	}
)
