package postNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"
	pb "microServiceBoilerplate/proto/generated/post"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type publishers struct {
	lgr *golog.Core
	nc  *nats.Conn
}

func (p *publishers) NewPost(userId, postId uint64) error {
	subject := configs.Nats.Subjects.NewPost
	dbug, sussess := p.lgr.DebugPKG("NewPost", false)

	{
		var (
			msgByte []byte

			err error
		)
		{
			msgByte, err = proto.Marshal(&natsPb.NewPost_EVENT{
				UserId: userId,
				PostId: postId,
			})
			if dbug("proto.Marshal")(err) != nil {
				return err
			}
		}
		{
			err = p.nc.Publish(subject, msgByte)
			if dbug("p.nc.Publish")(err) != nil {
				return err
			}
		}
	}
	sussess()
	return nil
}

func (p *publishers) GetUsers(userIds []uint64) (map[uint64]*pb.User, error) {
	subject := configs.Nats.Subjects.GetUsersByIds
	dbug, success := p.lgr.DebugPKG("GetProfiles", false)

	{
		var (
			request  []byte
			rawUsers *natsPb.GetUsers_REQUESTResponse
			result   map[uint64]*pb.User

			err error
		)

		{
			rawUsers = &natsPb.GetUsers_REQUESTResponse{}
		}
		{
			request, err = proto.Marshal(&natsPb.GetUsers_REQUESTRequest{
				UserIds: userIds,
			})
			if dbug("proto.Marshal")(err) != nil {
				return nil, err
			}
		}
		{
			var res *nats.Msg
			res, err := p.nc.Request(subject, request, configs.Nats.Timeout)
			if dbug("p.nc.Request")(err) != nil {
				return nil, err
			}
			err = proto.Unmarshal(res.Data, rawUsers)
			if dbug("proto.Unmarshal")(err) != nil {
				return nil, err
			}
		}
		{
			result = make(map[uint64]*pb.User, len(rawUsers.UsersData))
			for _, u := range rawUsers.UsersData {
				result[u.Id] = &pb.User{
					Name:   u.Name,
					Email:  u.Email,
					Id:     u.Id,
					Gender: u.Gender,
				}
			}
		}
		success("profiles: ", result)
		return result, nil
	}
}

func (p *publishers) IsFollowingGroup(userId uint64, followings []uint64) (map[uint64]bool, error) {
	subject := configs.Nats.Subjects.IsFollowingGroup
	dbug, success := p.lgr.DebugPKG("IsFollowings", false)

	{
		var (
			response     *natsPb.IsFollowingGroup_REQUESTResponse
			byteRequest  []byte
			byteResponse []byte

			err error
		)

		{
			response = &natsPb.IsFollowingGroup_REQUESTResponse{}
		}
		{
			byteRequest, err = proto.Marshal(&natsPb.IsFollowingGroup_REQUESTRequest{
				Follower:   userId,
				Followings: followings,
			})
			if dbug("proto.Marshal")(err) != nil {
				return nil, err
			}
		}
		{
			res, err := p.nc.Request(subject, byteRequest, configs.Nats.Timeout)
			if dbug("p.nc.Request")(err) != nil {
				return nil, err
			}
			byteResponse = res.Data
		}
		{
			err = proto.Unmarshal(byteResponse, response)
			if dbug("proto.Unmarshal")(err) != nil {
				return nil, err
			}
		}
		{
			success(response.Result)
			return response.Result, nil
		}
	}
}

func (p *publishers) GetCounts(postIds []uint64) (map[uint64]uint64, error) {
	subject := configs.Nats.Subjects.CountLikes
	dbug, success := p.lgr.DebugPKG("GetCounts", false)

	{
		var (
			byteRequest  []byte
			byteresponse []byte

			err error
		)

		{
			byteRequest, err = proto.Marshal(&natsPb.CountLikes_REQUESTRequest{
				PostId: postIds,
			})
			if dbug("proto.Marshal")(err) != nil {
				return nil, err
			}
		}
		{
			res, err := p.nc.Request(subject, byteRequest, configs.Nats.Timeout)
			if dbug("p.nc.Request")(err) != nil {
				return nil, err
			}
			byteresponse = res.Data
		}
		{
			response := &natsPb.CountLikes_REQUESTResponse{}
			err = proto.Unmarshal(byteresponse, response)
			if dbug("proto.Unmarshal")(err) != nil {
				return nil, err
			}
			success(response.Counts)
			return response.Counts, nil
		}
	}
}

func (p *publishers) IsLikedGroup(liker uint64, postIds []uint64) (map[uint64]*emptypb.Empty, error) {
	subject := configs.Nats.Subjects.IsLikedGroup
	dbug, sussess := p.lgr.DebugPKG("IsLikedGroup", false)

	{
		var (
			byteRequest  []byte
			byteResponse []byte

			err error
		)

		{
			byteRequest, err = proto.Marshal(&natsPb.IsLikedGroup_REQUESTRequest{
				PostIds: postIds,
				Liker:   liker,
			})
			if dbug("proto.Marshal")(err) != nil {
				return nil, err
			}
		}
		{
			res, err := p.nc.Request(subject, byteRequest, configs.Nats.Timeout)
			if dbug("p.nc.Request")(err) != nil {
				return nil, err
			}
			byteResponse = res.Data
		}
		{
			response := &natsPb.IsLikedGroup_REQUESTResponse{}
			err = proto.Unmarshal(byteResponse, response)
			if dbug("proto.Unmarshal")(err) != nil {
				return nil, err
			}
			sussess(response.Result)
			return response.Result, nil
		}
	}
}
