package repository

import (
	"github.com/go-redis/redis"
	"github.com/mreza0100/golog"
)

type read struct {
	lgr     *golog.Core
	db      *redis.Client
	helpers *helpers
}

func (r *read) GetFeed(userId, offset, limit uint64) ([]uint64, error) {
	vals := r.db.LRange(r.helpers.parseId(userId), int64(offset), int64(limit))
	if vals.Err() != nil {
		r.lgr.Debug.RedLog("error in GetFeed: ", vals.Err())
		return nil, vals.Err()
	}

	ids := make([]uint64, 0)
	vals.ScanSlice(&ids)

	return ids, nil
}
