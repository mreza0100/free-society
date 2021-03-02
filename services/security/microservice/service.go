package microservice

import (
	"microServiceBoilerplate/services/security/domain"
	"microServiceBoilerplate/services/security/handlers"
	"microServiceBoilerplate/services/security/instances"
	securityNats "microServiceBoilerplate/services/security/nats"

	"github.com/mreza0100/golog"
)

func NewSecurityService(lgr *golog.Core) instances.Handlers {
	nc := securityNats.Connection(lgr)

	publishers := securityNats.NewPublishers(&securityNats.NewPublishersOpts{
		Lgr: lgr,
		Nc:  nc,
	})

	services := domain.New(domain.ServiceOpts{
		Lgr: lgr,
	})

	securityNats.InitSubs(&securityNats.InitSubsOpts{
		Lgr: lgr,
		Srv: services,
		Nc:  nc,
	})

	return handlers.NewHandlers(handlers.NewHandlersOpts{
		Lgr:        lgr,
		Srv:        services,
		Publishers: publishers,
	})
}
