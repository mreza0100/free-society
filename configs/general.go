package configs

import (
	"fmt"
	"os"
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
	Token_expire                      = time.Hour * 24 * 7 // one week
	Token_expire_auto_remove_duration = time.Minute
)

var (
	// standard connection timeout for services
	stdConnectionTimeout = time.Duration(2 * time.Second)
	ROOT                 = os.Getenv("ROOT")
	LogPath              = ROOT + "/logs/all.log"
)

func str(thing interface{}) string {
	return fmt.Sprintf("%v", thing)
}
