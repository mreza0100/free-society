package domain

import (
	pb "freeSociety/proto/generated/user"
	"freeSociety/services/user/models"
	"freeSociety/utils/files/costume"
)

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
		AvatarPath:  costume.ExportAvatar(user.AvatarName),
		IsFollowing: isFollowing,
	}, nil
}
