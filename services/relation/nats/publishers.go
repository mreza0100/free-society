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
	subject := configs.NatsConfigs.Subjects.IsUserExist_REQUEST
	data := &natsPb.IsUserExist_REQUESTRequest{
		UserId: userId,
	}
	byteData, _ := proto.Marshal(data)

	{

		msg, _ := nc.Request(subject, byteData, configs.HellgateConfigs.Timeout)
		response := &natsPb.IsUserExist_REQUESTResponse{}
		proto.Unmarshal(msg.Data, response)

		this.lgr.Log("res: ", response.Exist)

		return response.Exist
	}
}
