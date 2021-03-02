package microservice

import (
	"microServiceBoilerplate/services/user/domain"
	"microServiceBoilerplate/services/user/handlers"
	"microServiceBoilerplate/services/user/instances"
	userNats "microServiceBoilerplate/services/user/nats"

	"github.com/mreza0100/golog"
)

func NewUserService(lgr *golog.Core) instances.Handlers {
	nc := userNats.Connection(lgr)

	services := domain.New(&domain.NewOpts{
		Lgr: lgr,
	})

	publishers := userNats.NewPublishers(nc, lgr)

	userNats.InitSubs(&userNats.InitSubsOpts{
		Lgr: lgr,
		Srv: services,
		Nc:  nc,
	})

	return handlers.NewHandlers(&handlers.NewHandlersOpts{
		Lgr:        lgr,
		Srv:        services,
		Publishers: publishers,
	})
}
