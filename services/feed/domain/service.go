package domain

import (
	"microServiceBoilerplate/services/feed/instances"
	"microServiceBoilerplate/services/feed/repository"

	"github.com/mreza0100/golog"
)

type NewOpts struct {
	Lgr        *golog.Core
	Publishers instances.Publishers
}

func New(opts *NewOpts) instances.Sevice {
	return &service{
		repo:       repository.New(opts.Lgr),
		lgr:        opts.Lgr.With("In domain->"),
		publishers: opts.Publishers,
	}
}

type service struct {
	repo       *instances.Repository
	lgr        *golog.Core
	publishers instances.Publishers
}

func (s *service) GetFeed(userId, offset, limit uint64) ([]uint64, error) {
	return s.repo.Read.GetFeed(userId, offset, limit)
}

func (s *service) SetPost(userId, postId uint64) error {
	followers, err := s.publishers.GetFollowers(userId)
	if err != nil {
		return err
	}

	return s.repo.Write.SetPostOnFeeds(userId, postId, followers)
}

func (s *service) DeleteFeed(userId uint64) error {
	return s.repo.Write.DeleteFeed(userId)
}
