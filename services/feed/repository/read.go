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

func (r *read) GetFeed(userId, offset, limit uint64) ([]string, error) {
	var (
		start int64
		stop  int64
		ids   []string
	)

	{
		start = int64(offset)
		stop = int64(offset + limit)
		ids = make([]string, 0, int(limit))
	}

	{
		vals := r.db.LRange(r.helpers.parseId(userId), start, stop)
		if vals.Err() != nil {
			r.lgr.Debug.RedLog("error in GetFeed: ", vals.Err())
			return nil, vals.Err()
		}

		vals.ScanSlice(&ids)
	}

	return ids, nil
}
