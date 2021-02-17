package microservice

import (
	"microServiceBoilerplate/services/feed/domain"
	"microServiceBoilerplate/services/feed/handlers"
	"microServiceBoilerplate/services/feed/instances"
	feedNats "microServiceBoilerplate/services/feed/nats"

	"github.com/mreza0100/golog"
)

func NewFeedService(lgr *golog.Core) instances.Handlers {
	var (
		services   instances.Sevice
		publishers instances.Publishers
		initSubs   func(instances.Sevice)
	)

	publishers, initSubs = feedNats.New(&feedNats.NewOpts{
		Lgr: lgr,
	})

	services = domain.New(&domain.NewOpts{
		Lgr:        lgr,
		Publishers: publishers,
	})

	defer initSubs(services)

	return handlers.New(&handlers.NewOpts{
		Lgr:        lgr,
		Srv:        services,
		Publishers: publishers,
	})
}
