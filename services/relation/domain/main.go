package domain

import (
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
