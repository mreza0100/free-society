package configs

import (
	"os"
	"time"
)

const (
	Token_expire                      = time.Hour * 24 * 7 // one week
	Token_expire_auto_remove_duration = time.Minute
	// 5 MB limit
	Picture_size_limit   = 5 * 1024 * 1024
	Max_picture_per_post = 4
)

var (
	ROOT         = os.Getenv("ROOT")
	LogPath      = ROOT + "/logs/all.log"
	PicturesPath = ROOT + "/public/images/"
)
