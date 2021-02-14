package feedNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"
	"microServiceBoilerplate/services/feed/types"

	"github.com/golang/protobuf/proto"
	"github.com/mreza0100/golog"
)

func NewPublishers(lgr *golog.Core) types.Publishers {
	return &publishers{
		lgr: lgr.With("In publishers: "),
	}
}

type publishers struct {
	lgr *golog.Core
}

func (this *publishers) GetFollowers(userId uint64) ([]uint64, error) {
	subject := configs.Nats.Subjects.GetFollowers

	{
		data := &natsPb.GetFollowers_REQUESTRequest{
			UserId: userId,
		}

		byteData, err := proto.Marshal(data)
		if err != nil {
			this.lgr.RedLog("In GetFollowers proto.Marshal error")
			this.lgr.RedLog("Error: ", err)
			return nil, err
		}
		response, err := nc.Request(subject, byteData, configs.Nats.Timeout)
		if err != nil {
			this.lgr.RedLog("In GetFollowers nc.Request error")
			this.lgr.RedLog("Error: ", err)
			return nil, err
		}

		followers := &natsPb.GetFollowers_REQUESTResponse{}

		err = proto.Unmarshal(response.Data, followers)
		if err != nil {
			this.lgr.RedLog("In GetFollowers proto.Unmarshal error")
			this.lgr.RedLog("Error: ", err)
			return nil, err
		}

		return followers.Followers, nil
	}
}

// not used
func (this *publishers) GetPosts(postIds []uint64) ([]*natsPb.Post, error) {
	subject := configs.Nats.Subjects.GetPosts

	{
		byteReq, err := proto.Marshal(&natsPb.GetPosts_REQUESTRequest{
			PostIds: postIds,
		})
		if err != nil {
			this.lgr.RedLog("In GetPosts proto.Marshal error")
			this.lgr.RedLog("Error: ", err)
			return nil, err
		}

		response, err := nc.Request(subject, byteReq, configs.Nats.Timeout)
		if err != nil {
			this.lgr.RedLog("In GetPosts nc.Request error")
			this.lgr.RedLog("Error: ", err)
			return nil, err
		}

		{
			data := &natsPb.GetPosts_REQUESTResponse{}
			err = proto.Unmarshal(response.Data, data)
			if err != nil {
				this.lgr.RedLog("In GetPosts proto.Unmarshal error")
				this.lgr.RedLog("Error: ", err)
				return nil, err
			}

			return data.Posts, nil
		}

	}
}
