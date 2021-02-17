package postNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"
	"microServiceBoilerplate/services/post/instances"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type initSubsOpts struct {
	lgr *golog.Core
	srv instances.Sevice
	nc  *nats.Conn
}

func initSubs(opts *initSubsOpts) {
	s := subscribers{
		srv: opts.srv,
		lgr: opts.lgr.With("In subscribers ->"),
	}
	opts.lgr.SuccessLog("subscribers has been attached to nats")

	s.DeleteUserPosts()
	s.GetPosts()
}

type subscribers struct {
	srv instances.Sevice
	lgr *golog.Core
	nc  *nats.Conn
}

func (s *subscribers) DeleteUserPosts() {
	subject := configs.Nats.Subjects.DeleteUser

	{
		s.nc.Subscribe(subject, func(msg *nats.Msg) {
			data := &natsPb.UserDelete_EVENT{}

			err := proto.Unmarshal(msg.Data, data)
			if err != nil {
				s.lgr.Debug.RedLog("In DeleteUserPosts: proto.Unmarshal has been returning error")
				s.lgr.Debug.RedLog("Error: ", err)
				return
			}
			err = s.srv.DeleteUserPosts(data.Id)
			if err != nil {
				s.lgr.Debug.RedLog("In DeleteUserPosts: DeleteUserPosts service has been returning error")
				s.lgr.Debug.RedLog("Error: ", err)
			}
		})
	}
}

func (s *subscribers) GetPosts() {
	subject := configs.Nats.Subjects.GetPosts

	{
		s.nc.Subscribe(subject, func(msg *nats.Msg) {
			data := &natsPb.GetPosts_REQUESTRequest{}

			err := proto.Unmarshal(msg.Data, data)
			if err != nil {
				s.lgr.Debug.RedLog("In GetPosts: proto.Unmarshal has been returning error")
				s.lgr.Debug.RedLog("Error: ", err)
				return
			}

			result, err := s.srv.GetPost(data.PostIds)
			if err != nil {
				s.lgr.Debug.RedLog("In GetPosts: service error")
				s.lgr.Debug.RedLog("Error: ", err)
				return
			}

			convertedResult := make([]*natsPb.Post, len(result))

			{
				// converting instances
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
					s.lgr.Debug.RedLog("In GetPosts: result Marshaling error")
					s.lgr.Debug.RedLog("Error: ", err)
					return
				}

				msg.Respond(dataBytes)
			}

		})
	}
}
