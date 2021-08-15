package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	securityPb "freeSociety/proto/generated/security"
	"freeSociety/proto/generated/user"
	models "freeSociety/services/hellgate/graph/model"
	"freeSociety/services/hellgate/security"
	"freeSociety/utils"
)

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (bool, error) {
	var (
		userId uint64
		token  string
	)

	{
		userRes, err := r.userConn.GetUser(ctx, &user.GetUserRequest{
			Email: email,
		})
		if err != nil {
			return false, errors.New("email or password is wrong")
		}
		userId = userRes.GetId()
	}

	{
		securityRes, err := r.SecurityConn.Login(ctx, &securityPb.LogInRequest{
			UserId:   userId,
			Password: password,
			Device:   security.GetUserAgent(ctx),
		})
		if err != nil {
			return false, utils.GetGRPCMSG(err)
		}
		token = securityRes.GetToken()
	}

	security.SetToken(ctx, token)

	return true, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	token := security.GetToken(ctx)

	r.SecurityConn.Logout(ctx, &securityPb.LogoutInRequest{
		Token: token,
	})

	security.DeleteToken(ctx)

	return true, nil
}

func (r *mutationResolver) DeleteSession(ctx context.Context, sessionID int) (bool, error) {
	_, err := r.SecurityConn.DeleteSession(ctx, &securityPb.DeleteSessionRequest{
		SessionId: uint64(sessionID),
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) ChangePassword(ctx context.Context, prevPassword string, newPassword string) (bool, error) {
	var (
		userId uint64
		err    error
	)

	{
		userId = security.GetUserId(ctx)
	}

	{
		_, err = r.SecurityConn.ChangePassword(ctx, &securityPb.ChangePasswordRequest{
			PrevPassword: prevPassword,
			NewPassword:  newPassword,
			UserId:       userId,
		})
	}

	return err == nil, err
}

func (r *queryResolver) Sessions(ctx context.Context) ([]*models.Session, error) {
	var (
		userId    uint64
		rawData   []*securityPb.Session
		converted []*models.Session
	)

	{
		userId = security.GetUserId(ctx)
	}
	{
		response, err := r.SecurityConn.GetSessions(ctx, &securityPb.GetSessionsRequest{
			UserId: userId,
		})
		if err != nil {
			return nil, err
		}
		rawData = response.Sessions
	}
	{
		converted = make([]*models.Session, len(rawData))
		for idx, i := range rawData {
			converted[idx] = &models.Session{
				SessionID: int(i.SessionId),
				CreatedAt: int(i.CreatedAt),
				Device:    i.Device,
			}
		}
	}

	return converted, nil
}
