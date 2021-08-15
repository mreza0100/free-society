package relationNats

import (
	"freeSociety/proto/connections"
	"freeSociety/services/relation/instances"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
)

const natName = "Relation Service"

func Connection(lgr *golog.Core) *nats.Conn {
	return connections.GetConnection(lgr, natName)
}

type NewPublishersOpts struct {
	Lgr *golog.Core
	Nc  *nats.Conn
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
	opts.Lgr.SuccessLog("subscribers has been attached to nats")

	s.GetFollowers()
	s.IsFollowingGroup()
	s.DeleteUser()
	s.isLikedGroup()
	s.countLikes()
}
func NewPublishers(nc *nats.Conn, lgr *golog.Core) instances.Publishers {
	return &publishers{
		lgr: lgr.With("In publishers->"),
		nc:  nc,
	}
}
