package domain

import (
	"fmt"
	pb "microServiceBoilerplate/proto/generated/user"
	"microServiceBoilerplate/services/user/models"

	"github.com/mreza0100/golog"
	"gorm.io/gorm"
)

type ServiceOptions struct {
	DB *gorm.DB
	Lg *golog.Core
}

func NewService(options ServiceOptions) Sevice {
	return &service{
		DB: options.DB,
		Lg: options.Lg,
	}
}

type Sevice interface {
	NewUser(in *pb.NewUserRequest) (uint64, error)
	GetUserById(id uint64) (*pb.GetUserByIdResponse, error)
	GetUsers() ([]*pb.UserinGetUsers, error)
	DeleteUserById(id uint64) error
}

type service struct {
	DB *gorm.DB
	Lg *golog.Core
	pb.UnimplementedUserServiceServer
}

func (this *service) NewUser(in *pb.NewUserRequest) (uint64, error) {
	user := models.User{
		Name:     in.Name,
		Gender:   in.Gender,
		Email:    in.Email,
		Password: in.Password,
	}

	this.Lg.Log("mamad")

	tx := this.DB.Create(&user)

	return user.ID, tx.Error
}

func (this *service) GetUserById(id uint64) (*pb.GetUserByIdResponse, error) {
	user := pb.GetUserByIdResponse{}

	tx := this.DB.Raw(`SELECT * FROM users WHERE id = ?`, id).Scan(&user)

	if tx.RowsAffected == -1 {
		return nil, fmt.Errorf("Not found")
	}

	return &user, tx.Error
}

func (this *service) GetUsers() ([]*pb.UserinGetUsers, error) {
	users := []*pb.UserinGetUsers{}

	tx := this.DB.Raw(`SELECT * FROM users`).Scan(&users)

	return users, tx.Error
}

func (this *service) DeleteUserById(id uint64) error {
	tx := this.DB.Exec(`DELETE FROM users WHERE id = ?`, id)

	if tx.RowsAffected == 0 {
		return fmt.Errorf("Not found")
	}

	return nil

}
