package resolvers

import (
	pbFeed "microServiceBoilerplate/proto/generated/feed"
	pbPost "microServiceBoilerplate/proto/generated/post"
	pbRelation "microServiceBoilerplate/proto/generated/relation"
	pbUser "microServiceBoilerplate/proto/generated/user"

	"microServiceBoilerplate/services/hellgate/graph/generated"
	gqlGenerated "microServiceBoilerplate/services/hellgate/graph/generated"
	"microServiceBoilerplate/services/hellgate/security"

	"github.com/mreza0100/golog"
)

type Resolver struct {
	Lgr *golog.Core

	feedConn     pbFeed.FeedServiceClient
	userConn     pbUser.UserServiceClient
	postConn     pbPost.PostServiceClient
	relationConn pbRelation.RelationServiceClient
}

type NewOpts struct {
	Lgr *golog.Core

	FeedConn     pbFeed.FeedServiceClient
	UserConn     pbUser.UserServiceClient
	PostConn     pbPost.PostServiceClient
	RelationConn pbRelation.RelationServiceClient
}

func New(opts NewOpts) *gqlGenerated.Config {
	resolvers := &Resolver{
		Lgr: opts.Lgr,

		feedConn:     opts.FeedConn,
		relationConn: opts.RelationConn,
		userConn:     opts.UserConn,
		postConn:     opts.PostConn,
	}

	directives := generated.DirectiveRoot{
		Private: security.PrivateRoute,
	}

	return &gqlGenerated.Config{
		Resolvers:  resolvers,
		Directives: directives,
	}
}
