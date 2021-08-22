package redis

import (
	"github.com/go-redis/redis"
	"github.com/mreza0100/golog"
)

type read struct {
	lgr *golog.Core
	db  *redis.Client
	h   *helpers
}

func (r *read) GetSession(token string) (uint64, error) {
	cmd := r.db.Get(r.h.addPrefixS(token))
	if cmd.Err() != nil {
		return 0, cmd.Err()
	}

	return cmd.Uint64()
}
