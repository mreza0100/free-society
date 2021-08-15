package microservice

import (
	"freeSociety/services/security/domain"
	"freeSociety/services/security/handlers"
	"freeSociety/services/security/instances"
	securityNats "freeSociety/services/security/nats"

	"github.com/mreza0100/golog"
)

func NewSecurityService(lgr *golog.Core) instances.Handlers {
	nc := securityNats.Connection(lgr)

	publishers := securityNats.NewPublishers(&securityNats.NewPublishersOpts{
		Lgr: lgr,
		Nc:  nc,
	})

	services := domain.New(&domain.NewOpts{
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
