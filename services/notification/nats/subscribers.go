package notificationNats

import (
	"freeSociety/configs"
	"freeSociety/services/notification/instances"

	natsPb "freeSociety/proto/generated/nats"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type subscribers struct {
	srv instances.Sevice
	lgr *golog.Core
	nc  *nats.Conn
}

func (s *subscribers) getLikeNotifications() {
	subject := configs.Nats.Subjects.SetLikeNotification
	dbug, success := s.lgr.DebugPKG("getLikeNotifications", false)

	{
		s.nc.Subscribe(subject, func(msg *nats.Msg) {
			var (
				request  = new(natsPb.SetNotificationRequest)
				response = new(natsPb.SetNotificationResponse)
			)

			{
				if dbug("proto.Unmarshal")(proto.Unmarshal(msg.Data, request)) != nil {
					return
				}
			}
			{
				var err error
				response.NotificationId, err = s.srv.SetLikeNotification(request.UserId, request.LikerId, request.PostId)
				if dbug("s.srv.SetLikeNotification")(err) != nil {
					return
				}
			}
			{
				byteResult, err := proto.Marshal(response)
				if dbug("proto.Marshal")(err) != nil {
					return
				}
				if dbug("msg.Respond")(msg.Respond(byteResult)) != nil {
					return
				}
			}
			success()
		})
	}
}
