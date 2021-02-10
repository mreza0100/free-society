package graph

import (
	pbPost "microServiceBoilerplate/proto/generated/post"
	pbUser "microServiceBoilerplate/proto/generated/user"
	gqlGenerated "microServiceBoilerplate/services/hellgate/graph/generated"

	"github.com/mreza0100/golog"
)

type Resolver struct {
	Lgr *golog.Core

	userConn pbUser.UserServiceClient
	postConn pbPost.PostServiceClient
}

type NewOpts struct {
	Lgr *golog.Core

	UserConn pbUser.UserServiceClient
	PostConn pbPost.PostServiceClient
}

func New(opts NewOpts) gqlGenerated.Config {
	var (
		resolvers gqlGenerated.ResolverRoot = &Resolver{
			Lgr: opts.Lgr,

			userConn: opts.UserConn,
			postConn: opts.PostConn,
		}
		directives gqlGenerated.DirectiveRoot = gqlGenerated.DirectiveRoot{}
	)

	return gqlGenerated.Config{
		Resolvers:  resolvers,
		Directives: directives,
	}
}
