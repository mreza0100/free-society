package microservice

import (
	"microServiceBoilerplate/services/relation/db"
	"microServiceBoilerplate/services/relation/domain"
	"microServiceBoilerplate/services/relation/handlers"
	relationNats "microServiceBoilerplate/services/relation/nats"
	"microServiceBoilerplate/services/relation/types"

	"github.com/mreza0100/golog"
)

func NewRelationService(lgr *golog.Core) types.Handlers {
	services := domain.NewService(domain.ServiceOptions{
		Lgr: lgr,
	})

	db.ConnectDB(lgr)
	relationNats.InitialNatsSubs(services, lgr)

	return handlers.NewHandlers(handlers.NewHandlersOpts{
		Srv:        services,
		Lgr:        lgr,
		Publishers: relationNats.NewPublishers(lgr),
	})
}
