package relationNats

import (
	"freeSociety/configs"
	natsPb "freeSociety/proto/generated/nats"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type publishers struct {
	lgr *golog.Core
	nc  *nats.Conn
}

func (p *publishers) LikeNotify(userId, likerId uint64, postId string) (uint64, error) {
	subject := configs.NatsConfigs.Subjects.SetLikeNotification
	dbug, success := p.lgr.DebugPKG("LikeNotify", false)

	{
		var (
			request      []byte
			byteResponse []byte
			response     = new(natsPb.SetNotificationResponse)
		)

		{
			var err error
			request, err = proto.Marshal(&natsPb.SetNotificationRequest{
				IsLike:  true,
				UserId:  userId,
				LikerId: likerId,
				PostId:  postId,
			})
			if dbug("proto.Marshal")(err) != nil {
				return 0, err
			}
		}
		{
			msg, err := p.nc.Request(subject, request, configs.NatsConfigs.Timeout)
			if dbug("p.nc.Request")(err) != nil {
				return 0, err
			}
			byteResponse = msg.Data
		}

		{
			err := proto.Unmarshal(byteResponse, response)
			if dbug("proto.Unmarshal")(err) != nil {
				return 0, err
			}
		}
		success()
		return response.NotificationId, nil
	}
}

func (p *publishers) IsPostsExists(postIds ...string) ([]string, error) {
	subject := configs.NatsConfigs.Subjects.IsPostsExists
	dbug, success := p.lgr.DebugPKG("IsPostsExists", false)

	{
		var (
			request      []byte
			byteResponse []byte
			response     *natsPb.IsPostsExists_REQUESTResponse

			err error
		)

		{
			request, err = proto.Marshal(&natsPb.IsPostsExists_REQUESTRequest{PostIds: postIds})
			if dbug("proto.Marshal")(err) != nil {
				return nil, err
			}
		}
		{
			response, err := p.nc.Request(subject, request, configs.NatsConfigs.Timeout)
			if dbug("p.nc.Request")(err) != nil {
				return nil, err
			}
			byteResponse = response.Data
		}
		{
			response = &natsPb.IsPostsExists_REQUESTResponse{}
			err = proto.Unmarshal(byteResponse, response)
			if dbug("proto.Unmarshal")(err) != nil {
				return nil, err
			}
		}

		success(response.Exists)
		return response.Exists, err
	}
}

func (p *publishers) IsUserExist(userId uint64) bool {
	subject := configs.NatsConfigs.Subjects.IsUserExist
	dbug, success := p.lgr.DebugPKG("IsUserExist", false)

	{
		byteData, _ := proto.Marshal(&natsPb.IsUserExist_REQUESTRequest{
			UserId: userId,
		})

		msg, err := p.nc.Request(subject, byteData, configs.NatsConfigs.Timeout)
		if dbug("nc.Request error")(err) != nil {
			return false
		}

		response := &natsPb.IsUserExist_REQUESTResponse{}
		err = proto.Unmarshal(msg.Data, response)
		if dbug("proto.Unmarshal error")(err) != nil {
			return false
		}

		success(response.Exist)
		return response.Exist
	}
}
