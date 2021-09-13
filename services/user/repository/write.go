package repository

import (
	dbHelper "freeSociety/utils/dbHelper"
	dbhelper "freeSociety/utils/dbHelper"

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

func (w *write) NewUser(name, gender, email, avatarName string) (id uint64, cc dbhelper.CommandController, err error) {
	if w.read.IsUserExistByEmail(email) {
		return 0, dbHelper.FakeCC(), status.Error(codes.AlreadyExists, "there is already a user with this email")
	}

	cc, err = dbHelper.Transaction(w.db, func(tx *gorm.DB) error {
		const query = `INSERT INTO users (name, gender, email, avatar_name) VALUES (?, ?, ?, ?) RETURNING id`
		params := []interface{}{name, gender, email, avatarName}

		tx = tx.Raw(query, params...)

		if tx.Error != nil {
			return tx.Error
		}
		result := new(struct{ Id uint64 })
		tx = tx.Scan(result)
		id = result.Id

		return nil
	})

	return id, cc, err
}

func (w *write) DeleteUser(userId uint64) (avatarName string, cc dbhelper.CommandController, err error) {
	cc, err = dbHelper.Transaction(w.db, func(tx *gorm.DB) error {
		const query = `DELETE FROM users WHERE id=? RETURNING avatar_name`
		params := []interface{}{userId}

		tx = tx.Raw(query, params...)

		result := new(struct{ AvatarName string })
		tx = tx.Scan(&result)
		avatarName = result.AvatarName

		return tx.Error
	})

	return avatarName, cc, err
}

func (w *write) UpdateUser(userId uint64, name, gender, avatarName string) (dbhelper.CommandController, error) {
	cc, err := dbHelper.Transaction(w.db, func(tx *gorm.DB) error {
		const query = `UPDATE users SET name = ?, gender = ?, avatar_name = ? WHERE id = ?`
		params := []interface{}{name, gender, avatarName, userId}

		tx = w.db.Exec(query, params...)
		return tx.Error
	})

	return cc, err
}
