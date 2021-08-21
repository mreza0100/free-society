package userNats

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

func (p *publishers) DeleteUser(userId uint64) error {
	var (
		byteData []byte
		err      error
	)

	{
		byteData, err = proto.Marshal(&natsPb.UserDelete_EVENT{
			Id: userId,
		})
		if err != nil {
			p.lgr.Log("cant marshal pb (DeleteUser)")
			p.lgr.Log("error: ", err)
			return err
		}
	}

	{
		err = p.nc.Publish(configs.NatsConfigs.Subjects.DeleteUser, byteData)
		if err != nil {
			p.lgr.Log("cant publish pb (DeleteUser)")
			p.lgr.Log("error: ", err)
		}
	}

	return err
}
func (p *publishers) IsFollowingGroup(userId uint64, followings []uint64) (map[uint64]bool, error) {
	subject := configs.NatsConfigs.Subjects.IsFollowingGroup
	dbug, success := p.lgr.DebugPKG("IsFollowings", false)

	{
		var (
			response     *natsPb.IsFollowingGroup_REQUESTResponse
			byteRequest  []byte
			byteResponse []byte

			err error
		)

		{
			response = &natsPb.IsFollowingGroup_REQUESTResponse{}
		}
		{
			byteRequest, err = proto.Marshal(&natsPb.IsFollowingGroup_REQUESTRequest{
				Follower:   userId,
				Followings: followings,
			})
			if dbug("proto.Marshal")(err) != nil {
				return nil, err
			}
		}
		{
			res, err := p.nc.Request(subject, byteRequest, configs.NatsConfigs.Timeout)
			if dbug("p.nc.Request")(err) != nil {
				return nil, err
			}
			byteResponse = res.Data
		}
		{
			err = proto.Unmarshal(byteResponse, response)
			if dbug("proto.Unmarshal")(err) != nil {
				return nil, err
			}
		}
		{
			success(response.Result)
			return response.Result, nil
		}
	}
}
