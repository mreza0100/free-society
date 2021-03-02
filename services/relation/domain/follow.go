package domain

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) Follow(follower, following uint64) error {
	if follower == following {
		return errors.New("you cant follow your self")
	}

	if s.repo.Followers_read.IsFollowing(follower, following) {
		return status.Error(codes.AlreadyExists, "already following")
	}

	return s.repo.Followers_write.SetFollower(follower, following)
}

func (s *service) Unfollow(following, follower uint64) error {
	return s.repo.Followers_write.RemoveFollow(following, follower)
}

func (s *service) GetFollowers(userId uint64) []uint64 {
	return s.repo.Followers_read.GetFollowers(userId)
}

func (s *service) DeleteAllRelations(userId uint64) error {
	return s.repo.Followers_write.DeleteAllRelations(userId)
}

func (s *service) IsFollowing(follower uint64, followings []uint64) (map[uint64]bool, error) {
	var (
		result   map[uint64]interface{}
		response map[uint64]bool

		err error
	)

	{
		result, err = s.repo.Followers_read.IsFollowingGroup(follower, followings)
		if err != nil {
			return nil, err
		}
	}
	{
		response = make(map[uint64]bool, len(followings))

		for _, i := range followings {
			_, isFollowing := result[i]
			response[i] = isFollowing
		}
	}

	return response, nil
}
