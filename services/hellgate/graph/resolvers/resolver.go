package resolvers

import (
	pbFeed "freeSociety/proto/generated/feed"
	pbNotification "freeSociety/proto/generated/notification"
	pbPost "freeSociety/proto/generated/post"
	pbRelation "freeSociety/proto/generated/relation"
	pbSecurity "freeSociety/proto/generated/security"
	pbUser "freeSociety/proto/generated/user"

	"freeSociety/services/hellgate/graph/generated"
	gqlGenerated "freeSociety/services/hellgate/graph/generated"
	"freeSociety/services/hellgate/security"

	"github.com/mreza0100/golog"
)

type Resolver struct {
	Lgr *golog.Core

	feedConn         pbFeed.FeedServiceClient
	userConn         pbUser.UserServiceClient
	postConn         pbPost.PostServiceClient
	SecurityConn     pbSecurity.SecurityServiceClient
	relationConn     pbRelation.RelationServiceClient
	notificationConn pbNotification.NotificationServiceClient
}

type NewOpts struct {
	Lgr *golog.Core

	FeedConn         pbFeed.FeedServiceClient
	UserConn         pbUser.UserServiceClient
	PostConn         pbPost.PostServiceClient
	SecurityConn     pbSecurity.SecurityServiceClient
	RelationConn     pbRelation.RelationServiceClient
	NotificationConn pbNotification.NotificationServiceClient
}

func New(opts NewOpts) *gqlGenerated.Config {
	resolvers := &Resolver{
		Lgr: opts.Lgr,

		feedConn:         opts.FeedConn,
		userConn:         opts.UserConn,
		postConn:         opts.PostConn,
		SecurityConn:     opts.SecurityConn,
		relationConn:     opts.RelationConn,
		notificationConn: opts.NotificationConn,
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
