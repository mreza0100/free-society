package db

import (
	"errors"
	fmt "fmt"
	pb "microServiceBoilerplate/proto/generated/user"
	"microServiceBoilerplate/services/user/models"

	"github.com/mreza0100/golog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DAOS struct {
	Lgr *golog.Core
}

func (this *DAOS) NewUser(user *models.User) (uint64, error) {
	const query = `INSERT INTO users (name, gender, email, password) VALUES (?, ?, ?, ?) RETURNING id`
	params := []interface{}{user.Name, user.Gender, user.Email, user.Password}

	tx := db.Raw(query, params...)
	if tx.Error != nil {
		return 0, tx.Error
	}

	userId := struct{ Id uint64 }{}
	tx = tx.Scan(&userId)

	return userId.Id, tx.Error

}
func (this *DAOS) GetUser(id uint64) (*pb.GetUserResponse, error) {
	const query = `SELECT * FROM users WHERE id=?`
	params := []interface{}{id}

	user := &pb.GetUserResponse{}
	tx := db.Raw(query, params...).Scan(user)

	if tx.RowsAffected != 1 {
		return nil, errors.New("Not found")
	}

	return user, nil

}

func (this *DAOS) DeleteUser(id uint64) error {
	const query = `DELETE FROM users WHERE id=?`
	params := []interface{}{id}

	tx := db.Exec(query, params...)

	if tx.RowsAffected == 0 {
		return status.Errorf(codes.NotFound, fmt.Sprintf("Not found %v", id))
	}

	return nil
}
func (this *DAOS) GetUserByEmail(email string) (*models.User, error) {
	const query = `SELECT * FROM users WHERE email=?`
	params := []interface{}{email}

	tx := db.Raw(query, params...)
	if tx.Error != nil {
		return nil, tx.Error
	}
	user := &models.User{}
	tx.Scan(user)

	return user, nil
}

func (this *DAOS) IsUserExist(userId uint64) bool {
	const query = `SELECT EXISTS(SELECT 1 FROM users WHERE id=?)`
	params := []interface{}{userId}

	tx := db.Raw(query, params...)
	if tx.Error != nil {
		return false
	}

	data := struct{ Exists bool }{}
	tx.Scan(&data)
	this.Lgr.Log(data)

	return data.Exists
}
