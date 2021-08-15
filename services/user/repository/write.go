package repository

import (
	models "freeSociety/services/user/models"

	"github.com/mreza0100/golog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	gorm "gorm.io/gorm"
)

type write struct {
	lgr  *golog.Core
	db   *gorm.DB
	read *read
}

func (w *write) NewUser(user *models.User) (uint64, error) {
	{
		if w.read.IsUserExistByEmail(user.Email) {
			const err = "there is already a user with this email"
			return 0, status.Error(codes.AlreadyExists, err)
		}
	}
	{
		const query = `INSERT INTO users (name, gender, email) VALUES (?, ?, ?) RETURNING id`
		params := []interface{}{user.Name, user.Gender, user.Email}

		tx := w.db.Raw(query, params...)
		if tx.Error != nil {
			return 0, tx.Error
		}

		data := struct{ Id uint64 }{}
		tx = tx.Scan(&data)

		return data.Id, tx.Error
	}
}

func (w *write) DeleteUser(userId uint64) error {
	const query = `DELETE FROM users WHERE id=?`
	params := []interface{}{userId}

	tx := w.db.Exec(query, params...)

	if tx.RowsAffected == 0 {
		return status.Error(codes.NotFound, "cant find user")
	}

	return nil
}
