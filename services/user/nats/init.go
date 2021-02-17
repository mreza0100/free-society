package userNats

import (
	"microServiceBoilerplate/proto/connections"
	"microServiceBoilerplate/services/user/instances"

	"github.com/mreza0100/golog"
)

const natName = "User Service"

type NewOpts struct {
	Lgr *golog.Core
	Srv instances.Sevice
}

func NewPublishers(opts *NewOpts) instances.Publishers {
	nc := connections.GetConnection(opts.Lgr, natName)

	{
		initSubs(&initSubsOpts{
			lgr: opts.Lgr.With("In Subscribers ->"),
			srv: opts.Srv,
			nc:  nc,
		})
	}

	return newPublishers(nc, opts.Lgr)
}
