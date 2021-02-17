package domain

import (
	"errors"
	"microServiceBoilerplate/services/relation/instances"
	"microServiceBoilerplate/services/relation/repository"

	"github.com/mreza0100/golog"
)

type NewOpts struct {
	Lgr *golog.Core
}

func New(opts *NewOpts) instances.Sevice {
	return &service{
		lgr:  opts.Lgr.With("In domain->"),
		repo: repository.NewRepo(opts.Lgr),
	}
}

type service struct {
	lgr  *golog.Core
	repo *instances.Repository
}

func (s *service) Follow(follower, following uint64) error {
	if follower == following {
		return errors.New("you cant follow your self")
	}

	return s.repo.Write.SetFollower(follower, following)
}

func (s *service) Unfollow(following, follower uint64) error {
	return s.repo.Write.RemoveFollow(following, follower)
}

func (s *service) GetFollowers(userId uint64) []uint64 {
	return s.repo.Read.GetFollowers(userId)
}

func (s *service) DeleteAllRelations(userId uint64) error {
	return s.repo.Write.DeleteAllRelations(userId)
}
