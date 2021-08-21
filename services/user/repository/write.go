package repository

import (
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

func (w *write) NewUser(name, gender, email, avatarPath string) (uint64, error) {
	{
		if w.read.IsUserExistByEmail(email) {
			return 0, status.Error(codes.AlreadyExists, "there is already a user with this email")
		}
	}
	{
		const query = `INSERT INTO users (name, gender, email, avatar_path) VALUES (?, ?, ?, ?) RETURNING id`
		params := []interface{}{name, gender, email, avatarPath}

		tx := w.db.Raw(query, params...)
		if tx.Error != nil {
			return 0, tx.Error
		}

		data := struct{ Id uint64 }{}
		tx = tx.Scan(&data)

		return data.Id, tx.Error
	}
}

func (w *write) DeleteUser(userId uint64) (picturePath string, err error) {
	const query = `DELETE FROM users WHERE id=? RETURNING avatar_path`
	params := []interface{}{userId}

	tx := w.db.Raw(query, params...)

	if tx.RowsAffected == 0 {
		return "", status.Error(codes.NotFound, "cant find user")
	}

	data := struct{ AvatarPath string }{}
	tx.Scan(&data)

	return data.AvatarPath, nil
}
