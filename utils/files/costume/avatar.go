package costume

import (
	"freeSociety/configs"
	"freeSociety/utils/files"
	"path"
)

func GetFullPathAvatar(name string) string {
	return path.Join(configs.ROOT, "/public/", configs.AvatarPath, name)
}

func SaveAvatar(name string, content []byte) error {
	return files.CreateAndWriteFile(GetFullPathAvatar(name), content)
}

func ExportAvatar(name string) string {
	return path.Join(configs.FilesDomain, configs.AvatarPath, name)
}

func DeleteAvatar(name string) error {
	return files.DeleteFile(GetFullPathAvatar(name))
}
