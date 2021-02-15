package random

import (
	rand "math/rand"
	time "time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetIntRange(r int) int {
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(r)))

	return rand.Intn(r)
}

func GetOneOfArray(arr []string) string {
	return arr[GetIntRange(len(arr)-1)]
}

func GetRandomStr(strLen int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, strLen)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
