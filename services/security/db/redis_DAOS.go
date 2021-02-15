package db

import (
	"microServiceBoilerplate/configs"

	"github.com/mreza0100/golog"
)

const SESSION_PREFIX = "sess:"

func (this *RedisDAOS) addPrefixS(token string) string {
	return SESSION_PREFIX + token
}
func (this *RedisDAOS) addPrefixes(tokens ...string) []string {
	for idx, t := range tokens {
		tokens[idx] = SESSION_PREFIX + t
	}
	return tokens
}

type RedisDAOS struct {
	Lgr *golog.Core
}

func (this *RedisDAOS) NewSession(token string, userId uint64) error {
	return redisDB.Set(this.addPrefixS(token), userId, configs.Token_expire).Err()
}

func (this *RedisDAOS) GetSession(token string) (uint64, error) {
	cmd := redisDB.Get(this.addPrefixS(token))
	if cmd.Err() != nil {
		this.Lgr.InfoLog(cmd.Err())
		return 0, cmd.Err()
	}

	return cmd.Uint64()
}

func (this *RedisDAOS) DeleteSession(token ...string) error {
	return redisDB.Del(this.addPrefixes(token...)...).Err()
}
