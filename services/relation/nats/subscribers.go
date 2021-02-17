package relationNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"
	"microServiceBoilerplate/services/relation/instances"

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
		nc:  opts.nc,
		lgr: opts.lgr.With("In subscribers->"),
	}
	opts.lgr.SuccessLog("subscribers has been attached to nats")

	s.GetFollowers_REQUEST()
}

type subscribers struct {
	srv instances.Sevice
	lgr *golog.Core
	nc  *nats.Conn
}

func (s *subscribers) GetFollowers_REQUEST() {
	subject := configs.Nats.Subjects.GetFollowers
	dbug, sussecc := s.lgr.DebugPKG("GetFollowers_REQUEST", false)

	{
		s.nc.Subscribe(subject, func(msg *nats.Msg) {
			var (
				userId    uint64
				response  []byte
				followers []uint64
				err       error
			)

			{
				request := &natsPb.GetFollowers_REQUESTRequest{}
				err = proto.Unmarshal(msg.Data, request)
				if dbug("cant Unmarshal request")(err) != nil {
					return
				}
				userId = request.GetUserId()
			}

			{
				followers = s.srv.GetFollowers(userId)
			}

			{
				response, err = proto.Marshal(&natsPb.GetFollowers_REQUESTResponse{
					Followers: followers,
				})
				if dbug("cant Marshal response")(err) != nil {
					return
				}
			}
			{
				sussecc(response)
				msg.Respond(response)
			}
		})
	}
}

func (s *subscribers) DeleteUser() {
	subject := configs.Nats.Subjects.DeleteUser
	debug, sussecc := s.lgr.DebugPKG("DeleteUser", false)

	{
		s.nc.Subscribe(subject, func(msg *nats.Msg) {
			var (
				userId uint64
				err    error
			)

			{
				request := &natsPb.UserDelete_EVENT{}
				err = proto.Unmarshal(msg.Data, request)
				if debug("proto.Unmarshal")(err) != nil {
					return
				}
				userId = request.GetId()
			}

			{
				sussecc(userId)
				err = s.srv.DeleteAllRelations(userId)
				if debug("s.srv.DeleteUser")(err) != nil {
					return
				}
			}
		})

	}
}
