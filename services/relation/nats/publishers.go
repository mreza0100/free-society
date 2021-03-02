package relationNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"
	"microServiceBoilerplate/services/relation/instances"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func newPublishers(nc *nats.Conn, lgr *golog.Core) instances.Publishers {
	publishers := publishers{
		lgr: lgr.With("In publishers->"),
		nc:  nc,
	}

	return &publishers
}

type publishers struct {
	lgr *golog.Core
	nc  *nats.Conn
}

func (p *publishers) IsPostsExists(postIds ...uint64) ([]uint64, error) {
	subject := configs.Nats.Subjects.IsPostsExists
	dbug, sussecc := p.lgr.DebugPKG("IsPostsExists", false)

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
			response, err := p.nc.Request(subject, request, configs.Nats.Timeout)
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

		sussecc(response.Exists)
		return response.Exists, err
	}
}

func (p *publishers) IsUserExist(userId uint64) bool {
	subject := configs.Nats.Subjects.IsUserExist_REQUEST
	dbug, sussecc := p.lgr.DebugPKG("IsUserExist", false)

	{
		byteData, _ := proto.Marshal(&natsPb.IsUserExist_REQUESTRequest{
			UserId: userId,
		})

		msg, err := p.nc.Request(subject, byteData, configs.HellgateConfigs.Timeout)
		if dbug("nc.Request error")(err) != nil {
			return false
		}

		response := &natsPb.IsUserExist_REQUESTResponse{}
		err = proto.Unmarshal(msg.Data, response)
		if dbug("proto.Unmarshal error")(err) != nil {
			return false
		}

		sussecc(response.Exist)
		return response.Exist
	}
}
