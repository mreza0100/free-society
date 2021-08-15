package microservice

import (
	"freeSociety/services/post/domain"
	"freeSociety/services/post/handlers"
	"freeSociety/services/post/instances"
	postNats "freeSociety/services/post/nats"

	"github.com/mreza0100/golog"
)

func NewPostService(lgr *golog.Core) instances.Handlers {
	nc := postNats.Connection(lgr)

	publishers := postNats.NewPublishers(nc, lgr)

	services := domain.New(&domain.NewOpts{
		Lgr:        lgr,
		Publishers: publishers,
	})

	postNats.InitSubs(&postNats.InitSubsOpts{
		Lgr: lgr,
		Srv: services,
		Nc:  nc,
	})

	return handlers.New(&handlers.NewOpts{
		Srv:        services,
		Publishers: publishers,
		Lgr:        lgr,
	})
}
