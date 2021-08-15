package microservice

import (
	"freeSociety/services/notification/domain"
	"freeSociety/services/notification/handlers"
	"freeSociety/services/notification/instances"
	nats "freeSociety/services/notification/nats"

	"github.com/mreza0100/golog"
)

func NewNotificationService(lgr *golog.Core) instances.Handlers {
	nc := nats.Connection(lgr)

	publishers := nats.NewPublishers(nc, lgr)

	services := domain.New(&domain.NewOpts{
		Lgr:        lgr,
		Publishers: publishers,
	})

	nats.InitSubs(&nats.InitSubsOpts{
		Lgr: lgr,
		Srv: services,
		Nc:  nc,
	})

	return handlers.New(&handlers.NewOpts{
		Srv:        services,
		Publishers: publishers,
		Lgr:        lgr,
	})
}
