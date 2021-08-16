package postNats

import (
	"freeSociety/configs"
	natsPb "freeSociety/proto/generated/nats"
	"freeSociety/services/post/instances"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type subscribers struct {
	srv instances.Sevice
	lgr *golog.Core
	nc  *nats.Conn
}

func (s *subscribers) DeleteUserPosts() {
	subject := configs.NatsConfigs.Subjects.DeleteUser
	dbug, success := s.lgr.DebugPKG("DeleteUserPosts", false)

	{
		s.nc.Subscribe(subject, func(msg *nats.Msg) {
			data := &natsPb.UserDelete_EVENT{}

			err := proto.Unmarshal(msg.Data, data)
			if dbug("proto.Unmarshal")(err) != nil {
				return
			}

			err = s.srv.DeleteUserPosts(data.Id)
			if dbug("s.srv.DeleteUserPosts")(err) != nil {
				return
			}
			success()
		})
	}
}

func (s *subscribers) IsExists() {
	subject := configs.NatsConfigs.Subjects.IsPostsExists
	dbug, success := s.lgr.DebugPKG("IsExists", false)

	{
		s.nc.Subscribe(subject, func(msg *nats.Msg) {
			var (
				request    *natsPb.IsPostsExists_REQUESTRequest
				response   *natsPb.IsPostsExists_REQUESTResponse
				lastResult []byte
				err        error
			)

			{
				request = &natsPb.IsPostsExists_REQUESTRequest{}
				err := proto.Unmarshal(msg.Data, request)
				if dbug("proto.Unmarshal")(err) != nil {
					return
				}
			}
			{
				result, err := s.srv.IsPostsExists(request.PostIds)
				if dbug("s.srv.IsPostsExists")(err) != nil {
					return
				}
				response = &natsPb.IsPostsExists_REQUESTResponse{}
				response.Exists = result
			}
			{
				lastResult, err = proto.Marshal(response)
				if dbug("proto.Marshal")(err) != nil {
					return
				}
			}

			msg.Respond(lastResult)
			success(response)
		})
	}
}

// ! not used
// func (s *subscribers) GetPosts() {
// 	subject := configs.Nats.Subjects.GetPosts

// 	{
// 		s.nc.Subscribe(subject, func(msg *nats.Msg) {
// 			data := &natsPb.GetPosts_REQUESTRequest{}

// 			err := proto.Unmarshal(msg.Data, data)
// 			if err != nil {
// 				s.lgr.Debug.RedLog("In GetPosts: proto.Unmarshal has been returning error")
// 				s.lgr.Debug.RedLog("Error: ", err)
// 				return
// 			}

// 			result, err := s.srv.GetPost(0, data.PostIds)
// 			if err != nil {
// 				s.lgr.Debug.RedLog("In GetPosts: service error")
// 				s.lgr.Debug.RedLog("Error: ", err)
// 				return
// 			}

// 			convertedResult := make([]*natsPb.Post, len(result))

// 			{
// 				// converting instances
// 				for idx, p := range result {
// 					convertedResult[idx] = &natsPb.Post{
// 						Title:   p.Title,
// 						Body:    p.Body,
// 						Id:      p.Id,
// 						OwnerId: p.OwnerId,
// 					}
// 				}
// 			}

// 			{
// 				dataBytes, err := proto.Marshal(&natsPb.GetPosts_REQUESTResponse{
// 					Posts: convertedResult,
// 				})
// 				if err != nil {
// 					s.lgr.Debug.RedLog("In GetPosts: result Marshaling error")
// 					s.lgr.Debug.RedLog("Error: ", err)
// 					return
// 				}

// 				msg.Respond(dataBytes)
// 			}

// 		})
// 	}
// }
