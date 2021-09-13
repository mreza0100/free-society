package instances

import (
	"freeSociety/services/user/models"
	dbhelper "freeSociety/utils/dbHelper"
)

type (
	read interface {
		GetUserById(id uint64) (*models.User, error)
		GetUserByEmail(email string) (*models.User, error)
		IsUserExistById(userId uint64) bool
		IsUserExistByEmail(email string) bool
		GetUsersByIds(userIds []uint64) ([]*models.User, error)
	}
	write interface {
		NewUser(name, gender, email, avatarName string) (uint64, dbhelper.CommandController, error)
		DeleteUser(userId uint64) (avatarName string, cc dbhelper.CommandController, err error)
		UpdateUser(userId uint64, name, gender, avatarName string) (dbhelper.CommandController, error)
	}

	Repository struct {
		Read  read
		Write write
	}
)
