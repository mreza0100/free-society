package repository

import (
	"errors"
	"microServiceBoilerplate/services/user/models"

	"github.com/mreza0100/golog"
	"gorm.io/gorm"
)

type read struct {
	lgr   *golog.Core
	db    *gorm.DB
	write *write
}

func (r *read) GetUserById(userId uint64) (*models.User, error) {
	const query = `SELECT * FROM users WHERE id=?`
	params := []interface{}{userId}

	user := &models.User{}
	tx := r.db.Raw(query, params...).Scan(user)

	if tx.RowsAffected != 1 {
		return nil, errors.New("Not found")
	}

	return user, nil

}

func (r *read) GetUserByEmail(email string) (*models.User, error) {
	const query = `SELECT * FROM users WHERE email=?`
	params := []interface{}{email}

	user := &models.User{}
	tx := r.db.Raw(query, params...).Scan(user)

	if tx.Error != nil || tx.RowsAffected == 0 {
		return nil, tx.Error
	}

	return user, nil
}

func (r *read) IsUserExistById(userId uint64) bool {
	const query = `SELECT EXISTS(SELECT 1 FROM users WHERE id=?)`
	params := []interface{}{userId}

	tx := r.db.Raw(query, params...)
	if tx.Error != nil {
		return false
	}

	data := struct{ Exists bool }{}
	tx.Scan(&data)

	return data.Exists
}

func (r *read) IsUserExistByEmail(email string) bool {
	const query = `SELECT EXISTS(SELECT 1 FROM users WHERE email=?)`
	params := []interface{}{email}

	tx := r.db.Raw(query, params...)
	if tx.Error != nil {
		return false
	}

	data := struct{ Exists bool }{}
	tx.Scan(&data)

	return data.Exists
}
