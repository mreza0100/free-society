package securityNats

import (
	"freeSociety/connections"
	"freeSociety/services/security/instances"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
)

const natName = "Security Service"

func initSubscribers(lgr *golog.Core, nc *nats.Conn, srv instances.Sevice) {
	s := subscribers{
		srv: srv,
		lgr: lgr.With("In subscribers->"),
		nc:  nc,
	}
	lgr.SuccessLog("subscribers has been attached to nats")

	{
		s.deleteDeletedUserSessions()
	}
}

func InitNats(lgr *golog.Core) (publishers instances.Publishers, setServices func(instances.Sevice)) {
	nc := connections.GetNatsConnection(lgr, natName)
	publishers = NewPublishers(nc, lgr)

	return publishers, func(s instances.Sevice) {
		initSubscribers(lgr, nc, s)
	}
}

func NewPublishers(nc *nats.Conn, lgr *golog.Core) instances.Publishers {
	return &publishers{
		lgr: lgr.With("In publishers->"),
		nc:  nc,
	}
}
