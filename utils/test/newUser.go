package test

import (
	"context"
	"freeSociety/proto/connections"
	"freeSociety/proto/generated/security"
	"freeSociety/proto/generated/user"

	"github.com/mreza0100/golog"
)

type NewUserOpns struct {
	Lgr    *golog.Core
	Name   string
	Email  string
	Gender string
}

func NewUser(opns NewUserOpns) (userId uint64, deleteUser func()) {
	userConn := connections.UserSrvConn(opns.Lgr)
	securityConn := connections.SecuritySrvConn(opns.Lgr)

	{
		response, err := userConn.NewUser(context.Background(), &user.NewUserRequest{
			Name:   opns.Name,
			Email:  opns.Email,
			Gender: opns.Gender,
		})
		if err != nil {
			panic(err)
		}
		if response.Id == 0 {
			panic("user id was 0")
		}
		userId = response.Id
	}
	{
		_, err := securityConn.NewUser(context.Background(), &security.NewUserRequest{
			UserId:   userId,
			Password: "8888888",
			Device:   "mamads device",
		})
		if err != nil {
			panic(err)
		}
	}

	deleteUser = func() {
		_, err := userConn.DeleteUser(context.Background(), &user.DeleteUserRequest{
			Id: userId,
		})
		if err != nil {
			panic(err)
		}
	}

	return
}
