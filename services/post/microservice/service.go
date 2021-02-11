package microservice

import (
	"microServiceBoilerplate/services/post/domain"
	"microServiceBoilerplate/services/post/handlers"
	postNats "microServiceBoilerplate/services/post/nats"

	"github.com/mreza0100/golog"
)

func NewPostService(Lgr *golog.Core) handlers.Handlers {
	services := domain.NewService(domain.ServiceOptions{
		Lgr: Lgr,
	})

	h := handlers.NewHandlers(services, Lgr)

	postNats.InitialNatsSubs(h, Lgr)
	return h
}
