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

func (this *read) GetSession(token string) (uint64, error) {
	cmd := this.db.Get(this.h.addPrefixS(token))
	if cmd.Err() != nil {
		this.lgr.InfoLog(cmd.Err())
		return 0, cmd.Err()
	}

	return cmd.Uint64()
}
