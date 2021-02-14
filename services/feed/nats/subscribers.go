package feedNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"
	"microServiceBoilerplate/services/feed/types"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func InitialNatsSubs(h types.Sevice, lgr *golog.Core) {
	subs := subscribers{
		srv: h,
		lgr: lgr,
	}

	lgr.GreenLog("✅ subscribers has been attached to nats")

	subs.SetPost()
}

type subscribers struct {
	srv types.Sevice
	lgr *golog.Core
}

func (this *subscribers) SetPost() {
	subject := configs.Nats.Subjects.NewPost

	{
		nc.Subscribe(subject, func(msg *nats.Msg) {
			data := &natsPb.NewPost_EVENT{}

			err := proto.Unmarshal(msg.Data, data)
			if err != nil {
				this.lgr.Debug.RedLog("proto.Unmarshal has been returning error")
				this.lgr.Debug.RedLog("Error: ", err)
			}

			this.srv.SetPost(data.UserId, data.PostId)
		})

	}
}
