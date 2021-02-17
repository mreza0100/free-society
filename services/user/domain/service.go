package domain

import (
	pb "microServiceBoilerplate/proto/generated/user"
	"microServiceBoilerplate/services/user/models"
	"microServiceBoilerplate/services/user/repository"

	"microServiceBoilerplate/services/user/instances"

	"github.com/mreza0100/golog"
)

type NewOpts struct {
	Lgr *golog.Core
}

func New(opts *NewOpts) instances.Sevice {
	return &service{
		repo: repository.NewRepo(opts.Lgr),
		lgr:  opts.Lgr.With("In domain->"),
	}
}

type service struct {
	repo *instances.Repository
	lgr  *golog.Core
}

func (s *service) NewUser(in *pb.NewUserRequest) (uint64, error) {
	user := models.User{
		Name:   in.Name,
		Gender: in.Gender,
		Email:  in.Email,
	}

	return s.repo.Write.NewUser(&user)
}

func (s *service) GetUser(id uint64, email string) (*pb.GetUserResponse, error) {
	var (
		user           *models.User
		err            error
		getQueryWithId = id != 0
	)

	{
		if getQueryWithId {
			user, err = s.repo.Read.GetUserById(id)
			if err != nil {
				return nil, err
			}
		} else {
			user, err = s.repo.Read.GetUserByEmail(email)
			if err != nil {
				return nil, err
			}
		}
	}

	return &pb.GetUserResponse{
		Id:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Gender: user.Gender,
	}, nil
}

func (s *service) DeleteUser(id uint64) error {
	return s.repo.Write.DeleteUser(id)
}

func (s *service) IsUserExist(userId uint64) bool {
	return s.repo.Read.IsUserExistById(userId)
}
