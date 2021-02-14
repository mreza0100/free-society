package microservice

import (
	"microServiceBoilerplate/services/security/db"
	"microServiceBoilerplate/services/security/domain"
	"microServiceBoilerplate/services/security/handlers"
	securityNats "microServiceBoilerplate/services/security/nats"
	"microServiceBoilerplate/services/security/types"

	"github.com/mreza0100/golog"
)

func NewUserService(lgr *golog.Core) types.Handlers {
	services := domain.NewService(domain.ServiceOptions{
		Lgr: lgr,
	})

	{
		db.ConnectPS(lgr)
		db.ConnecRedis(lgr)
	}

	securityNats.InitialNatsSubs(services, lgr)

	return handlers.NewHandlers(handlers.NewHandlersOpts{
		Lgr:        lgr,
		Srv:        services,
		Publishers: securityNats.NewPublishers(lgr),
	})
}
