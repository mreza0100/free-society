package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	securityPb "microServiceBoilerplate/proto/generated/security"
	"microServiceBoilerplate/proto/generated/user"
	"microServiceBoilerplate/services/hellgate/security"
	"microServiceBoilerplate/utils"
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
	token := security.GetCookie(ctx)

	r.SecurityConn.Logout(ctx, &securityPb.LogoutInRequest{
		Token: token,
	})

	security.DeleteToken(ctx)

	return true, nil
}
