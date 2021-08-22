package domain

import (
	"freeSociety/services/user/instances"
	"freeSociety/services/user/repository"

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
