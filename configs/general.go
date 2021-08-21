package configs

import (
	"freeSociety/utils/files"
	"os"
	"path"
	"time"
)

const (
	Token_expire                      = time.Hour * 24 * 7 // one week
	Token_expire_auto_remove_duration = time.Minute
	// 5 MB limit
	Picture_size_limit   = 1024 * 1024 * 5 // 5 MB
	Max_picture_per_post = 4

	DB_picture_sep  = ","
	Avatar_max_size = 1024 * 1024 * 5 // 5 MB
)

var (
	ROOT    = os.Getenv("ROOT")
	LogPath = ROOT + "/logs/all.log"

	PicturesPath = "/public/images/"
	AvatarPath   = "/public/avatars"

	FemaleDefaultAvatarPath = path.Join(AvatarPath, "default_female.jpg")
	MaleDefaultAvatarPath   = path.Join(AvatarPath, "default_male.jpeg")
)

func init() {
	if !files.FileExist(PicturesPath) {
		files.CreateDir(path.Join(PicturesPath))
	}

	if !files.FileExist(AvatarPath) {
		files.CreateDir(path.Join(AvatarPath))
	}
}
