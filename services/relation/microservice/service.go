package microservice

import (
	"microServiceBoilerplate/services/relation/domain"
	"microServiceBoilerplate/services/relation/handlers"
	"microServiceBoilerplate/services/relation/instances"
	relationNats "microServiceBoilerplate/services/relation/nats"

	"github.com/mreza0100/golog"
)

func NewRelationService(lgr *golog.Core) instances.Handlers {
	var (
		services   instances.Sevice
		publishers instances.Publishers
		initor     func(srv instances.Sevice)
	)

	{
		publishers, initor = relationNats.New(&relationNats.NewOpts{
			Lgr: lgr,
		})
	}

	{
		services = domain.New(&domain.NewOpts{
			Lgr: lgr,
		})
	}

	{
		initor(services)
	}

	return handlers.New(&handlers.NewOpts{
		Srv:        services,
		Publishers: publishers,
		Lgr:        lgr,
	})
}
