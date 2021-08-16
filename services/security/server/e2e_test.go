package main_test

import (
	"context"
	"errors"
	"freeSociety/connections"
	"freeSociety/proto/generated/security"
	"freeSociety/services/security/domain"
	"freeSociety/utils/test"
	"testing"

	"github.com/mreza0100/golog"
)

var (
	token = make(map[uint]string)
)

const (
	userId      = 85
	password    = "mamads_first_password"
	newPassword = "mamads_new_password"
)

var (
	logger = golog.New(golog.InitOpns{
		LogPath:   "",
		Name:      "security test",
		WithTime:  false,
		DebugMode: true,
	})
	securityConn = connections.SecuritySrvConn(logger)

	services = domain.New(&domain.NewOpts{
		Lgr: logger,
	})

	ctx = context.Background()
)

func TestMain(m *testing.M) {
	m.Run()
}

func Test_newUser(t *testing.T) {
	response, err := securityConn.NewUser(ctx, &security.NewUserRequest{
		UserId:   userId,
		Password: password,
		Device:   "mamads device",
	})
	test.CheckFail(t, err)
	test.FailIf(t, len(response.Token) != 40, errors.New("len(response.Token) != 40"))
	token[1] = response.Token
}

func Test_login(t *testing.T) {
	response, err := securityConn.Login(ctx, &security.LogInRequest{
		UserId:   userId,
		Password: password,
		Device:   "on yeki mamads device",
	})
	test.CheckFail(t, err)
	test.FailIf(t, len(response.Token) != 40, errors.New("len(response.Token) != 40"))
	token[2] = response.Token
}

func Test_sessions(t *testing.T) {
	response, err := securityConn.GetSessions(ctx, &security.GetSessionsRequest{
		UserId: userId,
	})
	test.CheckFail(t, err)
	test.FailIf(t, len(response.Sessions) != 2, errors.New("len(response.Sessions)!=2"))
}

func Test_getUserIdFromPassword(t *testing.T) {
	response, err := securityConn.GetUserId(ctx, &security.GetUserIdRequest{
		Token: token[1],
	})
	test.CheckFail(t, err)
	test.FailIf(t, response.UserId != userId, errors.New("response.UserId!=userId"))
}

func Test_logout(t *testing.T) {
	{
		_, err := securityConn.Logout(ctx, &security.LogoutInRequest{
			Token: token[1],
		})
		test.CheckFail(t, err)
	}
	{
		response, err := securityConn.GetSessions(ctx, &security.GetSessionsRequest{
			UserId: userId,
		})
		test.CheckFail(t, err)
		test.FailIf(t, len(response.Sessions) != 1, errors.New("len(response.Sessions)!=1"))
	}
}

func Test_deleteSession(t *testing.T) {
	var sessionId uint64

	{
		response, err := securityConn.GetSessions(ctx, &security.GetSessionsRequest{
			UserId: userId,
		})
		test.CheckFail(t, err)
		test.FailIf(t, len(response.Sessions) != 1, errors.New("len(response.Sessions)!=1"))

		sessionId = response.Sessions[0].SessionId
	}
	_, err := securityConn.DeleteSession(ctx, &security.DeleteSessionRequest{
		SessionId: sessionId,
	})
	test.CheckFail(t, err)
	{
		response, err := securityConn.GetSessions(ctx, &security.GetSessionsRequest{
			UserId: userId,
		})
		test.CheckFail(t, err)
		test.FailIf(t, len(response.Sessions) != 0, errors.New("len(response.Sessions)!=0"))
	}
}

func Test_changePassword(t *testing.T) {
	var newToken string

	{
		_, err := securityConn.ChangePassword(ctx, &security.ChangePasswordRequest{
			PrevPassword: password,
			NewPassword:  newPassword,
			UserId:       userId,
		})
		test.CheckFail(t, err)
	}
	{
		response, err := securityConn.Login(ctx, &security.LogInRequest{
			UserId:   userId,
			Password: newPassword,
			Device:   "some device",
		})
		test.CheckFail(t, err)
		newToken = response.Token
	}
	{
		response, err := securityConn.GetUserId(ctx, &security.GetUserIdRequest{
			Token: newToken,
		})
		test.CheckFail(t, err)
		test.FailIf(t, response.UserId != userId, errors.New("response.UserId!=userId"))
	}
}

func Test_deleteUser(t *testing.T) {
	err := services.PurgeUser(userId)
	test.CheckFail(t, err)
}
