package db

import (
	"microServiceBoilerplate/configs"

	"github.com/go-redis/redis"
	"github.com/mreza0100/golog"
)

var db *redis.Client

func ConnectDB(lgr *golog.Core) {
	db = redis.NewClient(&redis.Options{
		Addr:     "localhost:" + configs.FeedConfigs.StrDBPort,
		Password: "",
		DB:       0,

		OnConnect: func(c *redis.Conn) error {
			lgr.GreenLog("redis is connected successfuly")
			return nil
		},
	})
}
