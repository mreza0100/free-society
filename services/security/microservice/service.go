package microservice

import (
	pb "freeSociety/proto/generated/security"
	"freeSociety/services/security/domain"
	"freeSociety/services/security/handlers"
	securityNats "freeSociety/services/security/nats"

	"github.com/mreza0100/golog"
)

func NewSecurityService(lgr *golog.Core) pb.SecurityServiceServer {
	publishers, setServices := securityNats.InitNats(lgr)

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
