package microservice

import (
	"microServiceBoilerplate/services/user/domain"
	"microServiceBoilerplate/services/user/handlers"
	userNats "microServiceBoilerplate/services/user/nats"
	"microServiceBoilerplate/services/user/types"

	"github.com/mreza0100/golog"
)

func NewUserService(lgr *golog.Core) types.Handlers {
	services := domain.NewService(domain.ServiceOptions{
		Lgr: lgr,
	})

	h := handlers.NewHandlers(handlers.NewHandlersOpts{
		Srv:        services,
		Lgr:        lgr,
		Publishers: userNats.NewPublishers(lgr),
	})

	userNats.InitialNatsSubs(h, lgr)

	return h
}
