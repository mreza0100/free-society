package graph

import (
	"microServiceBoilerplate/proto/generated/user"
	"microServiceBoilerplate/services/hellgate/graph/generated"
)

type Resolver struct {
	userConn user.UserServiceClient
}

func New(userConn user.UserServiceClient) generated.Config {
	var (
		resolvers generated.ResolverRoot = &Resolver{
			userConn: userConn,
		}
		directives generated.DirectiveRoot = generated.DirectiveRoot{}
	)

	return generated.Config{
		Resolvers:  resolvers,
		Directives: directives,
	}
}
