package domain

import (
	"microServiceBoilerplate/services/relation/db"
	"microServiceBoilerplate/services/relation/types"

	"github.com/mreza0100/golog"
)

type ServiceOptions struct {
	Lgr *golog.Core
}

func NewService(opts ServiceOptions) types.Sevice {
	daos := db.DAOS{
		Lgr: opts.Lgr.With("in DAOS: "),
	}

	return &service{
		DAOS: daos,
		Lgr:  opts.Lgr.With("In domain: "),
	}
}

type service struct {
	DAOS db.DAOS
	Lgr  *golog.Core
}

func (s *service) Follow(following, follower uint64) error {
	return s.DAOS.SetFollower(following, follower)
}
