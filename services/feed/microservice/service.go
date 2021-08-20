package microservice

import (
	"freeSociety/services/feed/domain"
	"freeSociety/services/feed/handlers"
	"freeSociety/services/feed/instances"
	feedNats "freeSociety/services/feed/nats"

	"github.com/mreza0100/golog"
)

func NewFeedService(lgr *golog.Core) instances.Handlers {
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
