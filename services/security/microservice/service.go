package microservice

import (
	"freeSociety/services/security/domain"
	"freeSociety/services/security/handlers"
	"freeSociety/services/security/instances"
	securityNats "freeSociety/services/security/nats"

	"github.com/mreza0100/golog"
)

func NewSecurityService(lgr *golog.Core) instances.Handlers {
	publishers, setServices := securityNats.InitNats(lgr)

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
