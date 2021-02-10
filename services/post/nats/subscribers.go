package postNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"
	"microServiceBoilerplate/services/post/db"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type Nats struct {
	daos *db.DAOS
	lgr  *golog.Core
}

func InitialNatsSubs(daos *db.DAOS, lgr *golog.Core) {
	nats := Nats{
		daos: daos,
		lgr:  lgr,
	}
	lgr.Log("✅ subscribers has been attached to nats")

	nats.DeleteUserPosts()
}

func (this *Nats) DeleteUserPosts() {
	subject := configs.NatsConfigs.Subjects.DeleteUser

	{
		nc.Subscribe(subject, func(msg *nats.Msg) {
			data := natsPb.DeleteUserPosts_EVENT{}

			proto.Unmarshal(msg.Data, &data)

			err := this.daos.DeleteUserPosts(data.Id)
			if err != nil {
				this.lgr.Log("daos.DeleteUserPosts has been returning error")
				this.lgr.Log("error: ", err)
			}
		})
	}

}
