package main_test

import (
	"context"
	"errors"
	"freeSociety/connections"
	"freeSociety/proto/generated/security"
	"freeSociety/proto/generated/user"
	"freeSociety/utils/test"
	"testing"

	"github.com/mreza0100/golog"
)

const (
	name   = "mamad"
	email  = "mamad@mamad.com"
	gender = "male"
)

var (
	userId uint64
)

var (
	logger = golog.New(golog.InitOpns{
		LogPath:   "",
		Name:      "post test",
		WithTime:  false,
		DebugMode: true,
	})
	userConn     = connections.UserSrvConn(logger)
	securityConn = connections.SecuritySrvConn(logger)

	ctx = context.Background()
)

func Test_createUser(t *testing.T) {
	{
		response, err := userConn.NewUser(ctx, &user.NewUserRequest{
			Name:   name,
			Email:  email,
			Gender: gender,
		})
		test.CheckFail(t, err)
		userId = response.Id
	}
	{
		_, err := securityConn.NewUser(ctx, &security.NewUserRequest{
			UserId:   userId,
			Password: "88888888",
			Device:   "mamads device",
		})

		test.CheckFail(t, err)
	}
}

func Test_getUser(t *testing.T) {
	check := func(data *user.GetUserResponse) {
		test.FailIf(t, data.Name != name, errors.New("data.Name != name"))
		test.FailIf(t, data.Email != email, errors.New("data.Email!=email"))
		test.FailIf(t, data.Gender != gender, errors.New("data.Gender != gender"))
	}

	t.Run("get with email", func(t *testing.T) {
		response, err := userConn.GetUser(ctx, &user.GetUserRequest{
			Id:    0,
			Email: email,
		})
		test.CheckFail(t, err)
		check(response)
	})
	t.Run("get with id", func(t *testing.T) {
		response, err := userConn.GetUser(ctx, &user.GetUserRequest{
			Id:    userId,
			Email: "",
		})
		test.CheckFail(t, err)
		check(response)
	})
}

func Test_deleteuser(t *testing.T) {
	_, err := userConn.DeleteUser(ctx, &user.DeleteUserRequest{
		Id: userId,
	})
	test.CheckFail(t, err)
}
