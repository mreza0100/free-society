package microservice

import (
	"microServiceBoilerplate/proto/generated/user"
	"microServiceBoilerplate/services/user/domain"
	"microServiceBoilerplate/services/user/handlers"
	userNats "microServiceBoilerplate/services/user/nats"

	"github.com/mreza0100/golog"
)

func NewUserService(lgr *golog.Core) user.UserServiceServer {
	services := domain.NewService(domain.ServiceOptions{
		Lgr: lgr,
	})

	return handlers.NewHandlers(handlers.NewHandlersOpts{
		Srv:        services,
		Lgr:        lgr,
		Publishers: userNats.NewPublishers(lgr),
	})
}
