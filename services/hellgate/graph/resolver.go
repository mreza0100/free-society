package graph

import (
	"context"
	pbPost "microServiceBoilerplate/proto/generated/post"
	pbRelation "microServiceBoilerplate/proto/generated/relation"
	pbUser "microServiceBoilerplate/proto/generated/user"

	"microServiceBoilerplate/services/hellgate/graph/generated"
	gqlGenerated "microServiceBoilerplate/services/hellgate/graph/generated"
	"microServiceBoilerplate/services/hellgate/security"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mreza0100/golog"
)

type Resolver struct {
	Lgr *golog.Core

	userConn     pbUser.UserServiceClient
	postConn     pbPost.PostServiceClient
	RelationConn pbRelation.RelationServiceClient
}

type NewOpts struct {
	Lgr *golog.Core

	UserConn     pbUser.UserServiceClient
	PostConn     pbPost.PostServiceClient
	RelationConn pbRelation.RelationServiceClient
}

func privateRoute(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	CA := security.GetCookieAccess(ctx)
	if !CA.IsLoggedIn {
		return nil, CA.NotLoginErr
	}

	return next(ctx)
}

func New(opts NewOpts) *gqlGenerated.Config {
	resolvers := &Resolver{
		Lgr: opts.Lgr,

		RelationConn: opts.RelationConn,
		userConn:     opts.UserConn,
		postConn:     opts.PostConn,
	}

	directives := generated.DirectiveRoot{
		Private: privateRoute,
	}

	return &gqlGenerated.Config{
		Resolvers:  resolvers,
		Directives: directives,
	}
}
