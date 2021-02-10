package domain

import (
	"errors"
	pb "microServiceBoilerplate/proto/generated/user"
	"microServiceBoilerplate/services/user/db"
	"microServiceBoilerplate/services/user/models"
	"microServiceBoilerplate/utils/security"

	"github.com/mreza0100/golog"
)

type ServiceOptions struct {
	Lgr *golog.Core
}

func NewService(opts ServiceOptions) Sevice {
	daos := db.DAOS{
		Lgr: opts.Lgr.With("in DAOS: "),
	}

	return &service{
		DAOS: daos,
		Lgr:  opts.Lgr.With("In domain: "),
	}
}

type Sevice interface {
	NewUser(in *pb.NewUserRequest) (uint64, error)
	GetUser(id uint64) (*pb.GetUserResponse, error)
	DeleteUser(id uint64) error
	Validation(email, password string) (uint64, error)
}

type service struct {
	DAOS db.DAOS
	Lgr  *golog.Core
}

func (this *service) NewUser(in *pb.NewUserRequest) (uint64, error) {
	user := models.User{
		Name:     in.Name,
		Gender:   in.Gender,
		Email:    in.Email,
		Password: security.HashIt(in.Password),
	}

	return this.DAOS.NewUser(&user)
}

func (this *service) GetUser(id uint64) (*pb.GetUserResponse, error) {
	return this.DAOS.GetUser(id)
}

func (this *service) DeleteUser(id uint64) error {
	return this.DAOS.DeleteUser(id)
}

func (this *service) Validation(email, password string) (uint64, error) {
	user, err := this.DAOS.GetUserByEmail(email)
	if err != nil {
		return 0, errors.New("Password or email is wrong")
	}

	if !security.HashCompare(user.Password, password) {
		return 0, errors.New("Password or email is wrong")
	}

	return user.ID, nil
}
