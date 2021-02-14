package relationNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"
	"microServiceBoilerplate/services/relation/types"

	"github.com/mreza0100/golog"
	"google.golang.org/protobuf/proto"
)

func NewPublishers(lgr *golog.Core) types.Publishers {
	publishers := publishers{
		lgr: lgr.With("In publishers: "),
	}

	return &publishers
}

type publishers struct {
	lgr *golog.Core
}

func (this *publishers) IsUserExist(userId uint64) bool {
	subject := configs.Nats.Subjects.IsUserExist_REQUEST
	dbug := this.lgr.DebugPKG("IsUserExist")

	{
		byteData, _ := proto.Marshal(&natsPb.IsUserExist_REQUESTRequest{
			UserId: userId,
		})

		msg, err := nc.Request(subject, byteData, configs.HellgateConfigs.Timeout)
		if err != nil {
			dbug("nc.Request error")(err)
			return false
		}

		response := &natsPb.IsUserExist_REQUESTResponse{}
		err = proto.Unmarshal(msg.Data, response)
		if err != nil {
			dbug("proto.Unmarshal error")(err)
			return false
		}

		return response.Exist
	}
}
