package microservice

import (
	pb "freeSociety/proto/generated/feed"
	"freeSociety/services/feed/domain"
	"freeSociety/services/feed/handlers"
	feedNats "freeSociety/services/feed/nats"

	"github.com/mreza0100/golog"
)

func NewFeedService(lgr *golog.Core) pb.FeedServiceServer {
	publishers, setServices := feedNats.InitNats(lgr)

	services := domain.New(&domain.NewOpts{
		Lgr:        lgr,
		Publishers: publishers,
	})
	setServices(services)

	return handlers.New(&handlers.NewOpts{
		Lgr:        lgr,
		Srv:        services,
		Publishers: publishers,
	})
}
