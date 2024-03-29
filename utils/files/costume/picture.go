package costume

import (
	"freeSociety/configs"
	"freeSociety/utils/files"
	"path"
)

func GetFullPathPicture(name string) string {
	return path.Join(configs.ROOT, "/public/", configs.PicturesPath, name)
}

func SavePicture(name string, content []byte) error {
	return files.CreateAndWriteFile(GetFullPathPicture(name), content)
}

func ExportPicture(name string) string {
	return path.Join(configs.FilesDomain, configs.PicturesPath, name)
}

func DeletPicture(name string) error {
	return files.DeleteFile(GetFullPathPicture(name))
}
