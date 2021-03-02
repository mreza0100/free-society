package microservice

import (
	"microServiceBoilerplate/services/post/domain"
	"microServiceBoilerplate/services/post/handlers"
	"microServiceBoilerplate/services/post/instances"
	postNats "microServiceBoilerplate/services/post/nats"

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
	})

	return handlers.New(&handlers.NewOpts{
		Srv:        services,
		Publishers: publishers,
		Lgr:        lgr,
	})
}
