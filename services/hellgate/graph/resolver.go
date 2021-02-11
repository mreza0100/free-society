package graph

import (
	pbPost "microServiceBoilerplate/proto/generated/post"
	pbRelation "microServiceBoilerplate/proto/generated/relation"
	pbUser "microServiceBoilerplate/proto/generated/user"

	gqlGenerated "microServiceBoilerplate/services/hellgate/graph/generated"

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

func New(opts NewOpts) *gqlGenerated.Config {
	resolvers := &Resolver{
		Lgr: opts.Lgr,

		RelationConn: opts.RelationConn,
		userConn:     opts.UserConn,
		postConn:     opts.PostConn,
	}

	return &gqlGenerated.Config{
		Resolvers: resolvers,
	}
}
