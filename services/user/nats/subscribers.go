package userNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"
	"microServiceBoilerplate/services/user/instanses"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func InitialNatsSubs(srv instanses.Sevice, lgr *golog.Core) {
	s := subscribers{
		srv: srv,
		lgr: lgr.With("In subscribers => "),
	}
	lgr.GreenLog("✅ subscribers has been attached to nats")

	s.IsUserExist_REQUEST()
}

type subscribers struct {
	srv instanses.Sevice
	lgr *golog.Core
}

func (this *subscribers) IsUserExist_REQUEST() {
	subject := configs.Nats.Subjects.IsUserExist_REQUEST

	{
		nc.Subscribe(subject, func(msg *nats.Msg) {
			request := natsPb.IsUserExist_REQUESTRequest{}

			err := proto.Unmarshal(msg.Data, &request)
			if err != nil {
				this.lgr.RedLog("in Subscribe cant unMarshal request")
				this.lgr.RedLog("Error: ", err)
				return
			}

			exist := this.srv.IsUserExist(request.UserId)

			response := &natsPb.IsUserExist_REQUESTResponse{
				Exist: exist,
			}

			resByte, err := proto.Marshal(response)
			if err != nil {
				this.lgr.RedLog("in GetFollowers_REQUEST cant Marshal response")
				this.lgr.RedLog("Error: ", err)
				return
			}
			msg.Respond(resByte)
		})
	}
}
