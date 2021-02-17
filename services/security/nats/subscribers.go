package securityNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"
	"microServiceBoilerplate/services/security/instances"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type initSubsOpts struct {
	lgr *golog.Core
	srv instances.Sevice
	nc  *nats.Conn
}

func initSubs(opts *initSubsOpts) {
	s := subscribers{
		srv: opts.srv,
		nc:  opts.nc,
		lgr: opts.lgr.With("In subscribers->"),
	}
	defer opts.lgr.SuccessLog("subscribers has been attached to nats")

	s.deleteDeletedUserSessions()

}

type subscribers struct {
	srv instances.Sevice
	lgr *golog.Core
	nc  *nats.Conn
}

func (s *subscribers) deleteDeletedUserSessions() {
	subject := configs.Nats.Subjects.DeleteUser
	dbug, success := s.lgr.DebugPKG("deleteDeletedUserSessions", false)

	{
		s.nc.Subscribe(subject, func(msg *nats.Msg) {
			var (
				userId uint64
				err    error
			)

			{
				data := &natsPb.UserDelete_EVENT{}
				err = proto.Unmarshal(msg.Data, data)
				if dbug("proto.Unmarshal")(err) != nil {
					return
				}

				userId = data.Id
			}
			{
				err = s.srv.PurgeUser(userId)
				if dbug("s.srv.PurgeUser")(err) != nil {
					return
				}
			}
			success()
		})
	}
}
