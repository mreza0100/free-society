package notificationNats

import (
	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
)

type publishers struct {
	lgr *golog.Core
	nc  *nats.Conn
}
