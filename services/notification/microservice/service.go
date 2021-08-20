package microservice

import (
	"freeSociety/services/notification/domain"
	"freeSociety/services/notification/handlers"
	"freeSociety/services/notification/instances"
	notificationNats "freeSociety/services/notification/nats"

	"github.com/mreza0100/golog"
)

func NewNotificationService(lgr *golog.Core) instances.Handlers {
	publishers, setServices := notificationNats.InitNats(lgr)

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
