package notificationNats

import (
	"freeSociety/proto/connections"
	"freeSociety/services/notification/instances"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
)

const natName = "notification Service"

func Connection(lgr *golog.Core) *nats.Conn {
	return connections.GetConnection(lgr, natName)
}

type InitSubsOpts struct {
	Lgr *golog.Core
	Srv instances.Sevice
	Nc  *nats.Conn
}

func InitSubs(opts *InitSubsOpts) {
	s := subscribers{
		srv: opts.Srv,
		lgr: opts.Lgr.With("In subscribers->"),
		nc:  opts.Nc,
	}
	{
		s.getLikeNotifications()
	}
	opts.Lgr.SuccessLog("subscribers has been attached to nats")
}

func NewPublishers(nc *nats.Conn, lgr *golog.Core) instances.Publishers {
	return &publishers{
		lgr: lgr.With("In publishers->"),
		nc:  nc,
	}
}
