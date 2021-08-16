package securityNats

import (
	"freeSociety/configs"
	natsPb "freeSociety/proto/generated/nats"
	"freeSociety/services/security/instances"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type subscribers struct {
	srv instances.Sevice
	lgr *golog.Core
	nc  *nats.Conn
}

func (s *subscribers) deleteDeletedUserSessions() {
	subject := configs.NatsConfigs.Subjects.DeleteUser
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
