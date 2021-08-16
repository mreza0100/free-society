package configs

import (
	"os"
	"time"
)

const (
	Token_expire                      = time.Hour * 24 * 7 // one week
	Token_expire_auto_remove_duration = time.Minute
)

var (
	ROOT    = os.Getenv("ROOT")
	LogPath = ROOT + "/logs/all.log"
)
