package postNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"
	"microServiceBoilerplate/services/post/handlers"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type subscribers struct {
	handlers handlers.Handlers
	lgr      *golog.Core
}

func InitialNatsSubs(h handlers.Handlers, lgr *golog.Core) {
	s := subscribers{
		handlers: h,
		lgr:      lgr,
	}
	lgr.GreenLog("✅ subscribers has been attached to nats")

	s.DeleteUserPosts()
}

func (this *subscribers) DeleteUserPosts() {
	subject := configs.NatsConfigs.Subjects.DeleteUser

	{
		nc.Subscribe(subject, func(msg *nats.Msg) {
			data := natsPb.DeleteUserPosts_EVENT{}

			proto.Unmarshal(msg.Data, &data)

			err := this.handlers.DeleteUserPosts(data.Id)
			if err != nil {
				this.lgr.Debug.RedLog("daos.DeleteUserPosts has been returning error")
				this.lgr.Debug.RedLog("error: ", err)
			}
		})
	}
}
