package microservice

import (
	"microServiceBoilerplate/services/post/domain"
	"microServiceBoilerplate/services/post/handlers"
	"microServiceBoilerplate/services/post/instances"
	postNats "microServiceBoilerplate/services/post/nats"

	"github.com/mreza0100/golog"
)

func NewPostService(lgr *golog.Core) instances.Handlers {
	var (
		services   instances.Sevice
		publishers instances.Publishers
	)

	{
		services = domain.New(&domain.NewOpts{
			Lgr:        lgr,
			Publishers: publishers,
		})
	}
	{
		var initSubs func(instances.Sevice)
		publishers, initSubs = postNats.New(&postNats.NewOpts{
			Lgr: lgr,
		})
		initSubs(services)
	}

	return handlers.New(&handlers.NewOpts{
		Srv:        services,
		Publishers: publishers,
		Lgr:        lgr,
	})
}
