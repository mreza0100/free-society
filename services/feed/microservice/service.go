package microservice

import (
	"microServiceBoilerplate/services/feed/domain"
	"microServiceBoilerplate/services/feed/handlers"
	feedNats "microServiceBoilerplate/services/feed/nats"
	"microServiceBoilerplate/services/feed/types"

	"github.com/mreza0100/golog"
)

func NewFeedService(lgr *golog.Core) types.Handlers {
	publishers := feedNats.NewPublishers(lgr)

	services := domain.NewService(domain.NewSrvOpts{
		Lgr:        lgr,
		Publishers: publishers,
	})

	feedNats.InitialNatsSubs(services, lgr)

	return handlers.NewHandlers(&handlers.HandlersOpts{
		Srv:        services,
		Lgr:        lgr,
		Publishers: publishers,
	})
}
