package security

import (
	"crypto/sha1"
	"fmt"
	"os"
)

var (
	salt string
)

func init() {
	salt = os.Getenv("SALT")
	if salt == "" {
		panic("salt is not set in var envs")
	}
}

func HashIt(thing string) string {
	thing = salt + thing

	hash := sha1.New()

	hash.Write([]byte(thing))

	bs := hash.Sum(nil)
	hashedStr := fmt.Sprintf("%x", bs)

	return hashedStr
}

func HashCompare(hash, str string) bool {
	return hash == HashIt(str)
}
