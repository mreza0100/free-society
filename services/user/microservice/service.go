package microservice

import (
	"microServiceBoilerplate/proto/generated/user"
	"microServiceBoilerplate/services/user/db"
	"microServiceBoilerplate/services/user/domain"
	"microServiceBoilerplate/services/user/handlers"

	"github.com/mreza0100/golog"
)

func NewUserService(Lg *golog.Core) user.UserServiceServer {
	businessLogic := domain.NewService(domain.ServiceOptions{
		DB: db.DB,
		Lg: Lg.With("In domain: "),
	})

	return handlers.NewHandlers(businessLogic)
}
