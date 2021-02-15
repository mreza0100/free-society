package relationNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"
	"microServiceBoilerplate/services/relation/types"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func InitialNatsSubs(srv types.Sevice, lgr *golog.Core) {
	s := subscribers{
		srv: srv,
		lgr: lgr.With("In subscribers: "),
	}
	lgr.GreenLog("✅ subscribers has been attached to nats")

	s.GetFollowers_REQUEST()
}

type subscribers struct {
	srv types.Sevice
	lgr *golog.Core
}

func (this *subscribers) GetFollowers_REQUEST() {
	subject := configs.Nats.Subjects.GetFollowers
	dbug := this.lgr.DebugPKG("GetFollowers_REQUEST")

	{
		nc.Subscribe(subject, func(msg *nats.Msg) {
			request := natsPb.GetFollowers_REQUESTRequest{}

			err := proto.Unmarshal(msg.Data, &request)
			if err != nil {
				dbug("cant Unmarshal request")(err)
				return
			}

			followers := this.srv.GetFollowers(request.UserId)

			response := &natsPb.GetFollowers_REQUESTResponse{
				Followers: followers,
			}

			resByte, err := proto.Marshal(response)
			if err != nil {
				dbug("cant Marshal response")(err)
				return
			}
			msg.Respond(resByte)
		})
	}
}
