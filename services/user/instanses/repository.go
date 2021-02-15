package instanses

import "microServiceBoilerplate/services/user/models"

type (
	read interface {
		GetUserById(id uint64) (*models.User, error)
		GetUserByEmail(email string) (*models.User, error)
		IsUserExist(userId uint64) bool
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
