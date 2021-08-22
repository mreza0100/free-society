package domain

import (
	"freeSociety/services/security/instances"
	"freeSociety/services/security/repository/postgres"
	"freeSociety/services/security/repository/redis"

	"github.com/mreza0100/golog"
)

type NewOpts struct {
	Lgr        *golog.Core
	Publishers instances.Publishers
}

func New(opts *NewOpts) instances.Sevice {
	return &service{
		lgr:          opts.Lgr.With("In domain->"),
		redisRepo:    redis.New(opts.Lgr),
		postgresRepo: postgres.New(opts.Lgr),
	}
}

type service struct {
	redisRepo    *instances.Repo_Redis
	postgresRepo *instances.Repo_Postgres
	lgr          *golog.Core
}
