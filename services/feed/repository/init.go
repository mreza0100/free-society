package repository

import (
	"fmt"
	"microServiceBoilerplate/configs"
	"microServiceBoilerplate/services/feed/instances"
	"microServiceBoilerplate/utils"

	"github.com/go-redis/redis"
	"github.com/mreza0100/golog"
)

func New(lgr *golog.Core) *instances.Repository {
	var (
		connection *redis.Client
		h          *helpers

		readQ  *read
		writeQ *write
	)

	{
		connection = getConnection(lgr)
		h = &helpers{
			lgr: lgr,
		}
	}

	{
		readQ = &read{
			db:      connection,
			lgr:     lgr,
			helpers: h,
		}
		writeQ = &write{
			db:      connection,
			lgr:     lgr,
			helpers: h,
		}
	}

	return &instances.Repository{
		Read:  readQ,
		Write: writeQ,
	}
}

func getConnection(lgr *golog.Core) *redis.Client {
	connection := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%v", configs.FeedConfigs.RedisPort),
		Password: "",
		DB:       0,
	})
	if !utils.IsPong(connection.Ping()) {
		lgr.Fatal("ping was't pont")
	}
	lgr.SuccessLog("redis is connected successfuly")

	return connection
}
