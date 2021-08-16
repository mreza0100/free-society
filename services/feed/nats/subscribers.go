package feedNats

import (
	"freeSociety/configs"
	natsPb "freeSociety/proto/generated/nats"
	"freeSociety/services/feed/instances"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type subscribers struct {
	lgr *golog.Core
	srv instances.Sevice
	nc  *nats.Conn
}

func (s *subscribers) setPost() {
	subject := configs.NatsConfigs.Subjects.NewPost

	{
		s.nc.Subscribe(subject, func(msg *nats.Msg) {
			data := &natsPb.NewPost_EVENT{}

			err := proto.Unmarshal(msg.Data, data)
			if err != nil {
				s.lgr.Debug.RedLog("proto.Unmarshal has been returning error")
				s.lgr.Debug.RedLog("Error: ", err)
			}

			s.srv.SetPost(data.UserId, data.PostId)
		})

	}
}

func (s *subscribers) deleteFeed() {
	subject := configs.NatsConfigs.Subjects.DeleteUser
	dbug, success := s.lgr.DebugPKG("deleteFeed", false)

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
				err = s.srv.DeleteFeed(userId)
				dbug("s.srv.DeleteFeed")(err)
				success(userId)
			}
		})
	}
}
