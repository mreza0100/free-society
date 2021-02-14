package domain

import (
	"microServiceBoilerplate/services/feed/db"
	"microServiceBoilerplate/services/feed/types"

	"github.com/mreza0100/golog"
)

type NewSrvOpts struct {
	Lgr        *golog.Core
	Publishers types.Publishers
}

func NewService(opts NewSrvOpts) types.Sevice {
	db.ConnectDB(opts.Lgr.With("In db: "))
	daos := &db.DAOS{
		Lgr: opts.Lgr.With("In DAOS: "),
	}

	return &service{
		daos:       daos,
		lgr:        opts.Lgr.With("In domain: "),
		publishers: opts.Publishers,
	}
}

type service struct {
	daos       *db.DAOS
	lgr        *golog.Core
	publishers types.Publishers
}

func (s *service) GetFeed(userId, offset, limit uint64) ([]uint64, error) {
	return s.daos.GetFeed(userId, offset, limit)
}

func (s *service) SetPost(userId, postId uint64) error {
	followers, err := s.publishers.GetFollowers(userId)
	if err != nil {
		return err
	}

	return s.daos.SetPostOnFeeds(userId, postId, followers)
}
