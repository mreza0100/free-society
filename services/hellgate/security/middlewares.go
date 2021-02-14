package security

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
)

func PrivateRoute(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	r := GetRequest(ctx)

	userId, err := extractUserId(r)
	if err != nil {
		DeleteToken(ctx)
		return nil, errors.New("not login | why? : " + err.Error())
	}

	ctx = context.WithValue(ctx, UserIdKeyCtx, userId)

	return next(ctx)
}
