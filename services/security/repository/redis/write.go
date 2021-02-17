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

func (this *write) NewSession(token string, userId uint64) error {
	return this.db.Set(this.h.addPrefixS(token), userId, configs.Token_expire).Err()
}

func (this *write) DeleteSession(token ...string) error {
	return this.db.Del(this.h.addPrefixes(token...)...).Err()
}
