package userNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"
	"microServiceBoilerplate/services/user/types"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type subscribers struct {
	handlers types.Handlers
	lgr      *golog.Core
}

func InitialNatsSubs(handlers types.Handlers, lgr *golog.Core) {
	s := subscribers{
		handlers: handlers,
		lgr:      lgr.With("In subscribers => "),
	}
	lgr.GreenLog("✅ subscribers has been attached to nats")

	s.IsUserExist_REQUEST()
}

func (this *subscribers) IsUserExist_REQUEST() {
	subject := configs.NatsConfigs.Subjects.IsUserExist_REQUEST

	{
		nc.Subscribe(subject, func(msg *nats.Msg) {
			request := natsPb.IsUserExist_REQUESTRequest{}

			this.lgr.PurpleLog("im here")

			proto.Unmarshal(msg.Data, &request)

			exist := this.handlers.IsUserExist(request.UserId)

			response := &natsPb.IsUserExist_REQUESTResponse{
				Exist: exist,
			}

			resByte, _ := proto.Marshal(response)
			msg.Respond(resByte)
		})
	}
}
