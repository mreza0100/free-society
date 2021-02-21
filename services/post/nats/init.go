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
}

func New(opts *NewOpts) (instances.Publishers, func(srv instances.Sevice)) {
	var (
		nc *nats.Conn
	)

	{
		nc = connections.GetConnection(opts.Lgr, natName)
	}

	initSubs := func(srv instances.Sevice) {
		initSubs(&initSubsOpts{
			lgr: opts.Lgr.With("In Subscribers->"),
			srv: srv,
			nc:  nc,
		})
	}

	return newPublishers(nc, opts.Lgr), initSubs
}
