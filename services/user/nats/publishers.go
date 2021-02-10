package userNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"

	"github.com/mreza0100/golog"
	"google.golang.org/protobuf/proto"
)

type Publishers interface {
	DeleteUser(userId uint64) error
}

func NewPublishers(lgr *golog.Core) Publishers {
	publishers := publishersT{
		lgr: lgr.With("In publishers: "),
	}

	return &publishers
}

type publishersT struct {
	lgr *golog.Core
}

func (this *publishersT) DeleteUser(userId uint64) error {
	event := natsPb.DeleteUserPosts_EVENT{
		Id: userId,
	}
	data, err := proto.Marshal(&event)
	if err != nil {
		this.lgr.Log("cant marshal pb (DeleteUser)")
		this.lgr.Log("error: ", err)
		return err
	}

	err = nc.Publish(configs.NatsConfigs.Subjects.DeleteUser, data)

	if err != nil {
		this.lgr.Log("cant publish pb (DeleteUser)")
		this.lgr.Log("error: ", err)
	}

	return err
}