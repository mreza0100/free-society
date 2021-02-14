package microservice

import (
	"microServiceBoilerplate/services/post/domain"
	"microServiceBoilerplate/services/post/handlers"
	postNats "microServiceBoilerplate/services/post/nats"
	"microServiceBoilerplate/services/post/types"

	"github.com/mreza0100/golog"
)

func NewPostService(lgr *golog.Core) types.Handlers {
	services := domain.NewService(domain.ServiceOptions{
		Lgr: lgr,
	})
	postNats.InitialNatsSubs(services, lgr)

	return handlers.NewHandlers(&handlers.HandlersOptns{
		Srv:        services,
		Lgr:        lgr,
		Publishers: postNats.NewPublishers(lgr),
	})
}
