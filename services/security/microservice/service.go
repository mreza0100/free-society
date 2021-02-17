package microservice

import (
	"microServiceBoilerplate/services/security/domain"
	"microServiceBoilerplate/services/security/handlers"
	"microServiceBoilerplate/services/security/instances"
	securityNats "microServiceBoilerplate/services/security/nats"

	"github.com/mreza0100/golog"
)

func NewSecurityService(lgr *golog.Core) instances.Handlers {
	var (
		services   instances.Sevice
		publishers instances.Publishers
	)

	{
		services = domain.New(domain.ServiceOpts{
			Lgr: lgr,
		})

		publishers = securityNats.New(&securityNats.NewOpts{
			Lgr: lgr,
			Srv: services,
		})
	}

	return handlers.NewHandlers(handlers.NewHandlersOpts{
		Lgr:        lgr,
		Srv:        services,
		Publishers: publishers,
	})
}
