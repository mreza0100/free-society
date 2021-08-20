package microservice

import (
	"freeSociety/services/relation/domain"
	"freeSociety/services/relation/handlers"
	"freeSociety/services/relation/instances"
	relationNats "freeSociety/services/relation/nats"

	"github.com/mreza0100/golog"
)

func NewRelationService(lgr *golog.Core) instances.Handlers {
	publishers, setServices := relationNats.InitNats(lgr)

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
