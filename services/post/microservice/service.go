package microservice

import (
	pb "microServiceBoilerplate/proto/generated/post"
	"microServiceBoilerplate/services/post/db"
	"microServiceBoilerplate/services/post/domain"
	"microServiceBoilerplate/services/post/handlers"
	postNats "microServiceBoilerplate/services/post/nats"

	"github.com/mreza0100/golog"
)

func NewPostService(Lgr *golog.Core) pb.PostServiceServer {
	daos := &db.DAOS{
		Lgr: Lgr.With("In DAOS: "),
	}

	businessLogic := domain.NewService(domain.ServiceOptions{
		Lgr: Lgr,
	})
	postNats.InitialNatsSubs(daos, Lgr)

	return handlers.NewHandlers(businessLogic, Lgr)
}
