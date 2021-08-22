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
