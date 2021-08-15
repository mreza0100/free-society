package domain

import (
	"freeSociety/services/notification/instances"
	"freeSociety/services/notification/models"
	"freeSociety/services/notification/repository"

	"github.com/mreza0100/golog"
)

type NewOpts struct {
	Lgr        *golog.Core
	Publishers instances.Publishers
}

func New(opts *NewOpts) instances.Sevice {
	return &service{
		lgr:        opts.Lgr.With("In domain->"),
		repo:       repository.NewRepo(opts.Lgr),
		publishers: opts.Publishers,
	}
}

type service struct {
	lgr        *golog.Core
	repo       *instances.Repository
	publishers instances.Publishers
}

func (s *service) SetLikeNotification(userId, likerId, postId uint64) (uint64, error) {
	return s.repo.Write.SetLikeNotification(userId, likerId, postId)
}

func (s *service) GetNotifications(userId uint64, limit, offset int64) ([]models.Notification, error) {
	if limit > 50 {
		limit = 50
	}
	return s.repo.Read.GetNotifications(userId, limit, offset)
}

func (s *service) ClearNotifications(userId uint64) error {
	return s.repo.Write.ClearNotifications(userId)
}
