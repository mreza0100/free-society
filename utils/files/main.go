package files

import (
	"os"
	"strings"
)

// returns true if file exists
func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func CreateAndWriteFile(path string, content []byte) error {
	if FileExist(path) {
		return nil
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(content)
	return err
}

// extract format from file path
func GetFileFormat(path string) string {
	return path[strings.LastIndex(path, "."):]
}
