package microservice

import (
	pb "freeSociety/proto/generated/notification"
	"freeSociety/services/notification/domain"
	"freeSociety/services/notification/handlers"
	notificationNats "freeSociety/services/notification/nats"

	"github.com/mreza0100/golog"
)

func NewNotificationService(lgr *golog.Core) pb.NotificationServiceServer {
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
