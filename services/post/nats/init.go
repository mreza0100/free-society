package postNats

import (
	"microServiceBoilerplate/proto/connections"
	"microServiceBoilerplate/services/post/instances"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
)

const natName = "Post Service"

type NewOpts struct {
	Lgr *golog.Core
	Srv instances.Sevice
}

func New(opts *NewOpts) instances.Publishers {
	var (
		nc *nats.Conn
	)

	{
		nc = connections.GetConnection(opts.Lgr, natName)
	}

	{
		initSubs(&initSubsOpts{
			lgr: opts.Lgr.With("In Subscribers ->"),
			srv: opts.Srv,
			nc:  nc,
		})
	}

	return newPublishers(nc, opts.Lgr)
}
