package postNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"
	"microServiceBoilerplate/services/post/types"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func InitialNatsSubs(srv types.Sevice, lgr *golog.Core) {
	s := subscribers{
		srv: srv,
		lgr: lgr,
	}
	lgr.GreenLog("✅ subscribers has been attached to nats")

	s.DeleteUserPosts()
}

type subscribers struct {
	srv types.Sevice
	lgr *golog.Core
}

func (this *subscribers) DeleteUserPosts() {
	subject := configs.Nats.Subjects.DeleteUser

	{
		nc.Subscribe(subject, func(msg *nats.Msg) {
			data := &natsPb.DeleteUserPosts_EVENT{}

			err := proto.Unmarshal(msg.Data, data)
			if err != nil {
				this.lgr.Debug.RedLog("In DeleteUserPosts: proto.Unmarshal has been returning error")
				this.lgr.Debug.RedLog("Error: ", err)
				return
			}
			err = this.srv.DeleteUserPosts(data.Id)
			if err != nil {
				this.lgr.Debug.RedLog("In DeleteUserPosts: DeleteUserPosts service has been returning error")
				this.lgr.Debug.RedLog("Error: ", err)
			}
		})
	}
}

func (this *subscribers) GetPosts() {
	subject := configs.Nats.Subjects.GetPosts

	{
		nc.Subscribe(subject, func(msg *nats.Msg) {
			data := &natsPb.GetPosts_REQUESTRequest{}

			err := proto.Unmarshal(msg.Data, data)
			if err != nil {
				this.lgr.Debug.RedLog("In GetPosts: proto.Unmarshal has been returning error")
				this.lgr.Debug.RedLog("Error: ", err)
				return
			}

			result, err := this.srv.GetPost(data.PostIds)
			if err != nil {
				this.lgr.Debug.RedLog("In GetPosts: service error")
				this.lgr.Debug.RedLog("Error: ", err)
				return
			}

			convertedResult := make([]*natsPb.Post, len(result))

			{
				// converting types
				for idx, p := range result {
					convertedResult[idx] = &natsPb.Post{
						Title:   p.Title,
						Body:    p.Body,
						Id:      p.Id,
						OwnerId: p.OwnerId,
					}
				}
			}

			{
				dataBytes, err := proto.Marshal(&natsPb.GetPosts_REQUESTResponse{
					Posts: convertedResult,
				})
				if err != nil {
					this.lgr.Debug.RedLog("In GetPosts: result Marshaling error")
					this.lgr.Debug.RedLog("Error: ", err)
					return
				}

				msg.Respond(dataBytes)
			}

		})
	}
}
