package securityNats

import (
	"freeSociety/proto/connections"
	"freeSociety/services/security/instances"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
)

const natName = "Security Service"

func Connection(lgr *golog.Core) *nats.Conn {
	return connections.GetConnection(lgr, natName)
}

type NewPublishersOpts struct {
	Lgr *golog.Core
	Nc  *nats.Conn
}

func NewPublishers(opts *NewPublishersOpts) instances.Publishers {
	return newPublishers(opts.Nc, opts.Lgr)
}

type InitSubsOpts struct {
	Lgr *golog.Core
	Srv instances.Sevice
	Nc  *nats.Conn
}

func InitSubs(opts *InitSubsOpts) {
	s := subscribers{
		srv: opts.Srv,
		nc:  opts.Nc,
		lgr: opts.Lgr.With("In subscribers->"),
	}
	defer opts.Lgr.SuccessLog("subscribers has been attached to nats")

	s.deleteDeletedUserSessions()

}
