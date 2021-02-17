package securityNats

import (
	"microServiceBoilerplate/proto/connections"
	"microServiceBoilerplate/services/security/instances"

	"github.com/mreza0100/golog"
)

const natName = "Security Service"

type NewOpts struct {
	Lgr *golog.Core
	Srv instances.Sevice
}

func New(opts *NewOpts) instances.Publishers {
	nc := connections.GetConnection(opts.Lgr, natName)

	{
		initSubs(&initSubsOpts{
			lgr: opts.Lgr,
			srv: opts.Srv,
			nc:  nc,
		})
	}

	return newPublishers(nc, opts.Lgr)
}
