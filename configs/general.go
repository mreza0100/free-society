package configs

import (
	"fmt"
	"time"
)

type serviceConfigs struct {
	Addr    string
	Timeout time.Duration

	Port      int
	DBPort    int
	RedisPort int
}

const (
	LogPath = "./logs/all.log"
)

const (
	Token_expire = time.Hour * 24 * 7 // one week
)

var (
	// standard connection timeout for services
	stdConnectionTimeout = time.Duration(2 * time.Second)
)

func str(thing interface{}) string {
	return fmt.Sprintf("%v", thing)
}
