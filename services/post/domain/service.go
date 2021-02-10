package domain

import (
	"errors"
	pb "microServiceBoilerplate/proto/generated/post"
	"microServiceBoilerplate/services/post/db"

	"github.com/mreza0100/golog"
)

type Sevice interface {
	NewPost(title, body string, userId uint64) (uint64, error)
	GetPost(postIds []uint64) ([]*pb.GetPostResponseInner, error)
	DeletePost(postId, userId uint64) error
}

type ServiceOptions struct {
	Lgr  *golog.Core
	DAOS *db.DAOS
}

func NewService(opts ServiceOptions) Sevice {
	return &service{
		daos: opts.DAOS,
		lgr:  opts.Lgr.With("In domain: "),
	}
}

type service struct {
	daos *db.DAOS
	lgr  *golog.Core
}

func (this *service) NewPost(title, body string, userId uint64) (uint64, error) {
	return this.daos.NewPost(title, body, userId)
}

func (this *service) DeletePost(postId, userId uint64) error {
	return this.daos.DeletePost(postId, userId)
}

func (this *service) GetPost(postIds []uint64) ([]*pb.GetPostResponseInner, error) {
	if len(postIds) == 0 {
		return []*pb.GetPostResponseInner{}, errors.New("0 len in postIds")
	}
	return this.daos.GetPost(postIds)

}
