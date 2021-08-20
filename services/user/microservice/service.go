package microservice

import (
	"freeSociety/services/user/domain"
	"freeSociety/services/user/handlers"
	"freeSociety/services/user/instances"
	userNats "freeSociety/services/user/nats"

	"github.com/mreza0100/golog"
)

func NewUserService(lgr *golog.Core) instances.Handlers {
	publishers, setServices := userNats.InitNats(lgr)

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
