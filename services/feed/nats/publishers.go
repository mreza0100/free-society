package feedNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"

	"github.com/golang/protobuf/proto"
	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
)

type publishers struct {
	lgr *golog.Core
	nc  *nats.Conn
}

func (this *publishers) GetFollowers(userId uint64) ([]uint64, error) {
	subject := configs.Nats.Subjects.GetFollowers

	{
		var (
			byteRequest  []byte
			byteResponse []byte
			followers    []uint64
			err          error
		)
		{
			byteRequest, err = proto.Marshal(&natsPb.GetFollowers_REQUESTRequest{
				UserId: userId,
			})
			if err != nil {
				this.lgr.RedLog("In GetFollowers proto.Marshal error")
				this.lgr.RedLog("Error: ", err)
				return nil, err
			}
		}
		{
			response, err := this.nc.Request(subject, byteRequest, configs.Nats.Timeout)
			if err != nil {
				this.lgr.RedLog("In GetFollowers this.nc.Request error")
				this.lgr.RedLog("Error: ", err)
				return nil, err
			}
			byteResponse = response.Data
		}
		{
			response := &natsPb.GetFollowers_REQUESTResponse{}
			err = proto.Unmarshal(byteResponse, response)
			if err != nil {
				this.lgr.RedLog("In GetFollowers proto.Unmarshal error")
				this.lgr.RedLog("Error: ", err)
				return nil, err
			}
			followers = response.Followers
		}

		return followers, nil
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

		response, err := this.nc.Request(subject, byteReq, configs.Nats.Timeout)
		if err != nil {
			this.lgr.RedLog("In GetPosts this.nc.Request error")
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
