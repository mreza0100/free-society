package repository

import (
	"github.com/go-redis/redis"
	"github.com/mreza0100/golog"
)

type write struct {
	lgr     *golog.Core
	db      *redis.Client
	helpers *helpers
}

func (w *write) SetPostOnFeeds(userId uint64, postId string, followers []uint64) error {
	for _, follower := range followers {
		w.db.RPush(w.helpers.parseId(follower), postId)
	}

	return nil
}

func (w *write) DeleteFeed(userId uint64) error {
	cmd := w.db.Del(w.helpers.parseId(userId))
	return cmd.Err()
}
