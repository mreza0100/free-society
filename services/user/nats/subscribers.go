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
		lgr: opts.lgr.With("In subscribers->"),
		nc:  opts.nc,
	}
	opts.lgr.SuccessLog("subscribers has been attached to nats")

	s.isUserExist_REQUEST()
	s.getUsersByIds_REQUEST()
}

type subscribers struct {
	lgr *golog.Core
	srv instances.Sevice
	nc  *nats.Conn
}

func (s *subscribers) isUserExist_REQUEST() {
	subject := configs.Nats.Subjects.IsUserExist_REQUEST
	dbug, success := s.lgr.DebugPKG("IsUserExist_REQUEST", false)

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
				if dbug("Unmarshal request")(proto.Unmarshal(msg.Data, request)) != nil {
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
				if dbug("cant Marshal response")(err) != nil {
					return
				}
			}

			success()
			msg.Respond(byteResponse)
		})
	}
}

func (s *subscribers) getUsersByIds_REQUEST() {
	subject := configs.Nats.Subjects.GetUsersByIds
	dbug, success := s.lgr.DebugPKG("getUsersByIds_REQUEST", false)

	{
		s.nc.Subscribe(subject, func(msg *nats.Msg) {
			var (
				request  *natsPb.GetUsers_REQUESTRequest
				response *natsPb.GetUsers_REQUESTResponse
			)

			{
				request = &natsPb.GetUsers_REQUESTRequest{}
				response = &natsPb.GetUsers_REQUESTResponse{}
			}
			{
				if dbug("proto.Unmarshal")(proto.Unmarshal(msg.Data, request)) != nil {
					return
				}
			}
			{
				users, err := s.srv.GetUsers(request.UserIds)
				if dbug("s.srv.GetUsers")(err) != nil {
					return
				}
				response.UsersData = make(map[uint64]*natsPb.User, len(request.UserIds))
				for _, u := range users {
					response.UsersData[u.ID] = &natsPb.User{
						Name:   u.Name,
						Email:  u.Email,
						Id:     u.ID,
						Gender: u.Gender,
					}
				}
			}
			{
				byteResult, err := proto.Marshal(response)
				if dbug("proto.Marshal")(err) != nil {
					return
				}
				msg.Respond(byteResult)
				success(response)
			}
		})
	}
}
