package domain

import (
	"microServiceBoilerplate/services/security/db"
	"microServiceBoilerplate/services/security/types"

	"github.com/mreza0100/golog"
)

type ServiceOptions struct {
	Lgr *golog.Core
}

func NewService(opts ServiceOptions) types.Sevice {
	daos := db.PS_DAOS{
		Lgr: opts.Lgr.With("in Postgres DAOS: "),
	}

	return &service{
		PS_DAOS: daos,
		Lgr:     opts.Lgr.With("In domain: "),
	}
}

type service struct {
	PS_DAOS db.PS_DAOS
	Lgr     *golog.Core
}

func (s *service) NewUser(userId uint64, password string) (string, error) {
	return "", nil
}

func (s *service) Login(userId uint64, password string) (string, error) {
	return "", nil
}
