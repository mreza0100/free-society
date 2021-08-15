package microservice

import (
	"freeSociety/services/feed/domain"
	"freeSociety/services/feed/handlers"
	"freeSociety/services/feed/instances"
	feedNats "freeSociety/services/feed/nats"

	"github.com/mreza0100/golog"
)

func NewFeedService(lgr *golog.Core) instances.Handlers {
	nc := feedNats.Connection(lgr)

	publishers := feedNats.NewPublishers(nc, lgr)

	service := domain.New(&domain.NewOpts{
		Lgr:        lgr,
		Publishers: publishers,
	})

	feedNats.InitSubs(&feedNats.InitSubsOpts{
		Lgr: lgr,
		Srv: service,
		Nc:  nc,
	})

	return handlers.New(&handlers.NewOpts{
		Lgr:        lgr,
		Srv:        service,
		Publishers: publishers,
	})
}
