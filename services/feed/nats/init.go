package feedNats

import (
	"microServiceBoilerplate/proto/connections"
	"microServiceBoilerplate/services/feed/instances"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
)

const natName = "Feed Service"

type NewOpts struct {
	Lgr *golog.Core
}

func New(opts *NewOpts) (instances.Publishers, func(srv instances.Sevice)) {
	var (
		nc     *nats.Conn
		initor func(srv instances.Sevice)
	)

	{
		nc = connections.GetConnection(opts.Lgr, natName)
	}

	{
		initor = func(srv instances.Sevice) {
			initSubs(&initSubsOpts{
				lgr: opts.Lgr,
				srv: srv,
				nc:  nc,
			})
		}
	}

	return newPublishers(nc, opts.Lgr), initor
}
