package securityNats

import (
	"freeSociety/services/security/instances"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
)

func newPublishers(nc *nats.Conn, lgr *golog.Core) instances.Publishers {
	return &publishers{
		lgr: lgr.With("In publishers->"),
	}
}

type publishers struct {
	lgr *golog.Core
}
