package utils

import (
	"fmt"
	"microServiceBoilerplate/utils/random"
	"microServiceBoilerplate/utils/security"

	uu "github.com/gofrs/uuid"
)

func Getuuid() string {
	id, _ := uu.NewV4()
	return id.String()
}

func CreateToken() string {
	uuid := Getuuid()
	uuid += fmt.Sprintf("%v", random.GetIntRange(1000))
	return security.HashIt(uuid)
}
