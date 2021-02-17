package microservice

import (
	"microServiceBoilerplate/services/user/domain"
	"microServiceBoilerplate/services/user/handlers"
	"microServiceBoilerplate/services/user/instances"
	userNats "microServiceBoilerplate/services/user/nats"

	"github.com/mreza0100/golog"
)

func NewUserService(lgr *golog.Core) instances.Handlers {
	var (
		publishers instances.Publishers
		services   instances.Sevice
	)

	{
		services = domain.New(&domain.NewOpts{
			Lgr: lgr,
		})
	}

	{
		publishers = userNats.NewPublishers(&userNats.NewOpts{
			Lgr: lgr,
			Srv: services,
		})
	}

	return handlers.NewHandlers(&handlers.NewHandlersOpts{
		Lgr:        lgr,
		Srv:        services,
		Publishers: publishers,
	})
}
