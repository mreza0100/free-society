package hash

import (
	"crypto/sha512"
	"encoding/hex"
)

func Hash512(data []byte) string {
	hasher := sha512.New()
	hasher.Write(data)
	byteResult := hasher.Sum(nil)

	return hex.EncodeToString(byteResult)
}
