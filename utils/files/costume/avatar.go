package costume

import (
	"freeSociety/configs"
	"freeSociety/utils/files"
	"path"
)

func SaveAvatar(name string, content []byte) error {
	p := path.Join(configs.ROOT, "/public/", configs.AvatarPath, name)
	return files.CreateAndWriteFile(p, content)
}

func ExportAvatar(name string) string {
	return path.Join(configs.FilesDomain, configs.AvatarPath, name)
}

func DeletAvatar(name string) error {
	p := path.Join(configs.ROOT, "/public/", configs.AvatarPath, name)
	return files.DeleteFile(p)
}
