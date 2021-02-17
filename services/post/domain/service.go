package domain

import (
	pb "microServiceBoilerplate/proto/generated/post"
	"microServiceBoilerplate/services/post/instances"
	"microServiceBoilerplate/services/post/repository"

	"github.com/mreza0100/golog"
)

type NewOpts struct {
	Lgr *golog.Core
}

func New(opts *NewOpts) instances.Sevice {
	return &service{
		lgr:  opts.Lgr.With("In domain ->"),
		repo: repository.NewRepo(opts.Lgr),
	}
}

type service struct {
	lgr  *golog.Core
	repo *instances.Repository
}

func (this *service) NewPost(title, body string, userId uint64) (uint64, error) {
	return this.repo.Write.NewPost(title, body, userId)
}

func (this *service) DeletePost(postId, userId uint64) error {
	return this.repo.Write.DeletePost(postId, userId)
}

func (this *service) GetPost(postIds []uint64) ([]*pb.Post, error) {
	if len(postIds) == 0 {
		return []*pb.Post{}, nil
	}
	return this.repo.Read.GetPost(postIds)

}
func (this *service) DeleteUserPosts(userId uint64) error {
	return this.repo.Write.DeleteUserPosts(userId)
}
