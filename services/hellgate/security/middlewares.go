package security

import (
	"context"
	"errors"
	"microServiceBoilerplate/proto/generated/security"

	"github.com/99designs/gqlgen/graphql"
)

func PrivateRoute(securityConn security.SecurityServiceClient) func(context.Context, interface{}, graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		var (
			token  string
			userId uint64
		)

		{
			req := GetRequest(ctx)
			auth, err := req.Cookie(COOKIE_NAME)
			if err != nil {
				return nil, errors.New("U are not logged in")
			}
			token = auth.Value
		}

		{
			response, err := securityConn.GetUserId(ctx, &security.GetUserIdRequest{
				Token: token,
			})
			if err != nil {
				DeleteToken(ctx)
				return nil, errors.New("U are not logged in")
			}
			userId = response.UserId
		}

		return next(context.WithValue(ctx, USER_ID_KEY_CTX, userId))
	}
}

func OptinalRoute(securityConn security.SecurityServiceClient) func(context.Context, interface{}, graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		var (
			token  string
			userId uint64
		)

		{
			req := GetRequest(ctx)
			auth, err := req.Cookie(COOKIE_NAME)
			if err != nil {
				return next(ctx)
			}
			token = auth.Value
		}

		{
			response, err := securityConn.GetUserId(ctx, &security.GetUserIdRequest{
				Token: token,
			})
			if err != nil {
				DeleteToken(ctx)
				return next(ctx)
			}
			userId = response.UserId
		}

		return next(context.WithValue(ctx, USER_ID_KEY_CTX, userId))
	}
}
