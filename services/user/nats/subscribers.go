package userNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"
	"microServiceBoilerplate/services/user/instances"

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
		nc:  opts.nc,
	}
	opts.lgr.SuccessLog("subscribers has been attached to nats")

	s.IsUserExist_REQUEST()
}

type subscribers struct {
	lgr *golog.Core
	srv instances.Sevice
	nc  *nats.Conn
}

func (s *subscribers) IsUserExist_REQUEST() {
	subject := configs.Nats.Subjects.IsUserExist_REQUEST

	{
		s.nc.Subscribe(subject, func(msg *nats.Msg) {
			var (
				request  *natsPb.IsUserExist_REQUESTRequest
				response *natsPb.IsUserExist_REQUESTResponse

				isExist      bool
				byteResponse []byte

				err error
			)
			{
				request = &natsPb.IsUserExist_REQUESTRequest{}
				response = &natsPb.IsUserExist_REQUESTResponse{}
			}

			{
				if proto.Unmarshal(msg.Data, request) != nil {
					s.lgr.RedLog("in Subscribe cant unMarshal request")
					s.lgr.RedLog("Error: ", err)
					return
				}
			}

			{
				isExist = s.srv.IsUserExist(request.UserId)
				response = &natsPb.IsUserExist_REQUESTResponse{
					Exist: isExist,
				}
			}

			{
				byteResponse, err = proto.Marshal(response)
				if err != nil {
					s.lgr.RedLog("in GetFollowers_REQUEST cant Marshal response")
					s.lgr.RedLog("Error: ", err)
					return
				}
			}

			msg.Respond(byteResponse)
		})
	}
}
