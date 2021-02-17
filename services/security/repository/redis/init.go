package redis

import (
	"fmt"
	"microServiceBoilerplate/configs"
	"microServiceBoilerplate/services/security/instances"
	"microServiceBoilerplate/utils"

	"github.com/go-redis/redis"
	"github.com/mreza0100/golog"
)

const (
	SESSION_PREFIX = "sess:"
)

func New(lgr *golog.Core) *instances.Repo_Redis {
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
			db:  connection,
			lgr: lgr,
			h:   h,
		}
		writeQ = &write{
			db:  connection,
			lgr: lgr,
			h:   h,
		}
	}

	return &instances.Repo_Redis{
		Read:  readQ,
		Write: writeQ,
	}
}

func getConnection(lgr *golog.Core) *redis.Client {
	connection := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%v", configs.SecurityConfigs.RedisPort),
		Password: "",
		DB:       0,
	})
	if !utils.IsPong(connection.Ping()) {
		lgr.Fatal("ping was't pont")
	}
	lgr.SuccessLog("redis is connected successfuly")

	return connection
}
