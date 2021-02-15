package repository

import (
	fmt "fmt"
	models "microServiceBoilerplate/services/user/models"

	"github.com/mreza0100/golog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	gorm "gorm.io/gorm"
)

type write struct {
	lgr *golog.Core
	db  *gorm.DB
}

func (w *write) NewUser(user *models.User) (uint64, error) {
	const query = `INSERT INTO users (name, gender, email) VALUES (?, ?, ?) RETURNING id`
	params := []interface{}{user.Name, user.Gender, user.Email}

	tx := w.db.Raw(query, params...)
	if tx.Error != nil {
		return 0, tx.Error
	}

	userId := struct{ Id uint64 }{}
	tx = tx.Scan(&userId)

	return userId.Id, tx.Error

}

func (w *write) DeleteUser(id uint64) error {
	const query = `DELETE FROM users WHERE id=?`
	params := []interface{}{id}

	tx := w.db.Exec(query, params...)

	if tx.RowsAffected == 0 {
		return status.Errorf(codes.NotFound, fmt.Sprintf("Not found %v", id))
	}

	return nil
}
