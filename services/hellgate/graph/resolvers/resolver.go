package resolvers

import (
	pbFeed "microServiceBoilerplate/proto/generated/feed"
	pbPost "microServiceBoilerplate/proto/generated/post"
	pbRelation "microServiceBoilerplate/proto/generated/relation"
	pbSecurity "microServiceBoilerplate/proto/generated/security"
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
	SecurityConn pbSecurity.SecurityServiceClient
	relationConn pbRelation.RelationServiceClient
}

type NewOpts struct {
	Lgr *golog.Core

	FeedConn     pbFeed.FeedServiceClient
	UserConn     pbUser.UserServiceClient
	PostConn     pbPost.PostServiceClient
	SecurityConn pbSecurity.SecurityServiceClient
	RelationConn pbRelation.RelationServiceClient
}

func New(opts NewOpts) *gqlGenerated.Config {
	resolvers := &Resolver{
		Lgr: opts.Lgr,

		feedConn:     opts.FeedConn,
		userConn:     opts.UserConn,
		postConn:     opts.PostConn,
		SecurityConn: opts.SecurityConn,
		relationConn: opts.RelationConn,
	}

	directives := generated.DirectiveRoot{
		// I need Security connection in my middleware
		Private:  security.PrivateRoute(opts.SecurityConn),
		Optional: security.OptinalRoute(opts.SecurityConn),
	}

	return &gqlGenerated.Config{
		Resolvers:  resolvers,
		Directives: directives,
	}
}
