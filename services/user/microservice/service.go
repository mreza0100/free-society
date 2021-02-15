package microservice

import (
	"microServiceBoilerplate/services/user/domain"
	"microServiceBoilerplate/services/user/handlers"
	userNats "microServiceBoilerplate/services/user/nats"

	"microServiceBoilerplate/services/user/instanses"

	"github.com/mreza0100/golog"
)

func NewUserService(lgr *golog.Core) instanses.Handlers {
	services := domain.NewService(domain.ServiceOpts{
		Lgr: lgr,
	})

	userNats.InitialNatsSubs(services, lgr)

	return handlers.NewHandlers(handlers.NewHandlersOpts{
		Srv:        services,
		Lgr:        lgr,
		Publishers: userNats.NewPublishers(lgr),
	})
}
