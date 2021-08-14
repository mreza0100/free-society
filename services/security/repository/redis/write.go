package redis

import (
	"microServiceBoilerplate/configs"

	"github.com/go-redis/redis"
	"github.com/mreza0100/golog"
)

type write struct {
	lgr *golog.Core
	db  *redis.Client
	h   *helpers
}

func (w *write) NewSession(token string, userId uint64) error {
	return w.db.Set(w.h.addPrefixS(token), userId, configs.Token_expire).Err()
}

func (w *write) DeleteSession(tokens ...string) error {
	if len(tokens) == 0 {
		return nil
	}
	return w.db.Del(w.h.addPrefixes(tokens...)...).Err()
}
