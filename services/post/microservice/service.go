package microservice

import (
	"freeSociety/services/post/domain"
	"freeSociety/services/post/handlers"
	"freeSociety/services/post/instances"
	postNats "freeSociety/services/post/nats"

	"github.com/mreza0100/golog"
)

func NewPostService(lgr *golog.Core) instances.Handlers {
	publishers, setServices := postNats.InitNats(lgr)

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
