package postNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"
	"microServiceBoilerplate/services/post/types"

	"github.com/mreza0100/golog"
	"google.golang.org/protobuf/proto"
)

func NewPublishers(lgr *golog.Core) types.Publishers {
	return &publishers{
		lgr: lgr.With("In publishers: "),
	}
}

type publishers struct {
	lgr *golog.Core
}

func (this *publishers) NewPost(userId, postId uint64) error {
	subject := configs.Nats.Subjects.NewPost

	{
		data := &natsPb.NewPost_EVENT{
			UserId: userId,
			PostId: postId,
		}
		msgByte, err := proto.Marshal(data)
		if err != nil {
			this.lgr.Debug.RedLog("proto.Marshal has been returning error")
			this.lgr.Debug.RedLog("Error: ", err)
			return err
		}

		err = nc.Publish(subject, msgByte)
		if err != nil {
			this.lgr.RedLog("in NewPost: can't publish msg")
			this.lgr.RedLog("error: ", err)
			return err
		}
	}
	return nil
}
