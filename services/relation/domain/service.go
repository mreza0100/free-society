package domain

import (
	"errors"
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

func (s *service) Follow(follower, following uint64) error {
	if follower == following {
		return errors.New("you cant follow your self")
	}

	return s.DAOS.SetFollower(follower, following)
}

func (s *service) Unfollow(following, follower uint64) error {
	return s.DAOS.RemoveFollower(following, follower)
}
func (this *service) GetFollowers(userId uint64) []uint64 {
	return this.DAOS.GetFollowers(userId)
}
