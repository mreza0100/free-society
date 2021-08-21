package domain

import (
	"freeSociety/configs"
	pb "freeSociety/proto/generated/user"
	"freeSociety/services/user/models"
	"freeSociety/services/user/repository"
	"freeSociety/utils/files"
	"path"

	"freeSociety/services/user/instances"

	"github.com/gofrs/uuid"
	"github.com/mreza0100/golog"
)

type NewOpts struct {
	Lgr        *golog.Core
	Publishers instances.Publishers
}

func New(opts *NewOpts) instances.Sevice {
	return &service{
		repo:       repository.NewRepo(opts.Lgr),
		lgr:        opts.Lgr.With("In domain->"),
		publishers: opts.Publishers,
	}
}

type service struct {
	repo       *instances.Repository
	lgr        *golog.Core
	publishers instances.Publishers
}

func (s *service) NewUser(name, email, gender, avatarFormat string, avatar []byte) (uint64, error) {
	var (
		avatarName string
	)

	{
		if len(avatar) == 0 {
			if gender == "male" {
				avatarName = configs.MaleDefaultAvatarPath
			} else {
				avatarName = configs.FemaleDefaultAvatarPath
			}
		} else {
			id, err := uuid.NewV4()
			if err != nil {
				return 0, err
			}
			avatarName = id.String() + avatarFormat

			err = files.CreateAndWriteFile(path.Join(configs.ROOT, configs.AvatarPath, avatarName), avatar)
			if err != nil {
				return 0, err
			}
		}
	}

	return s.repo.Write.NewUser(name, gender, email, avatarName)
}

func (s *service) GetUser(requestorId, id uint64, email string) (*pb.GetUserResponse, error) {
	var (
		user           *models.User
		isFollowing    bool
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
	{
		isFollowingGroup, err := s.publishers.IsFollowingGroup(requestorId, []uint64{user.ID})
		if err != nil {
			return nil, err
		}
		isFollowing = isFollowingGroup[user.ID]
	}

	return &pb.GetUserResponse{
		Id:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Gender:      user.Gender,
		AvatarPath:  user.AvatarPath,
		IsFollowing: isFollowing,
	}, nil
}

func (s *service) DeleteUser(id uint64) error {
	avatarName, err := s.repo.Write.DeleteUser(id)
	if err != nil {
		return err
	}

	if avatarName == configs.MaleDefaultAvatarPath || avatarName == configs.FemaleDefaultAvatarPath {
		return nil
	}

	return files.DeleteFile(path.Join(configs.ROOT, configs.AvatarPath, avatarName))
}

func (s *service) IsUserExist(userId uint64) bool {
	return s.repo.Read.IsUserExistById(userId)
}

func (s *service) GetUsers(ids []uint64) (map[uint64]*models.User, error) {
	var (
		rawUsers []*models.User
		result   map[uint64]*models.User
		err      error
	)

	{
		rawUsers, err = s.repo.Read.GetUsersByIds(ids)
		if err != nil {
			return nil, err
		}
	}

	{
		result = make(map[uint64]*models.User, len(rawUsers))
		for _, u := range rawUsers {
			result[u.ID] = &models.User{
				Name:       u.Name,
				Email:      u.Email,
				ID:         u.ID,
				Gender:     u.Gender,
				AvatarPath: u.AvatarPath,
				CreatedAt:  u.CreatedAt,
			}
		}
	}

	return result, err
}
