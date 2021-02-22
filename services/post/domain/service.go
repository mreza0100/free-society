package domain

import (
	pb "microServiceBoilerplate/proto/generated/post"
	"microServiceBoilerplate/services/post/instances"
	"microServiceBoilerplate/services/post/repository"
	"microServiceBoilerplate/utils"

	"github.com/mreza0100/golog"
)

type NewOpts struct {
	Lgr        *golog.Core
	Publishers instances.Publishers
}

func New(opts *NewOpts) instances.Sevice {
	return &service{
		lgr:        opts.Lgr.With("In domain->"),
		repo:       repository.NewRepo(opts.Lgr),
		publishers: opts.Publishers,
	}
}

type service struct {
	lgr        *golog.Core
	repo       *instances.Repository
	publishers instances.Publishers
}

func (s *service) NewPost(title, body string, userId uint64) (uint64, error) {
	return s.repo.Write.NewPost(title, body, userId)
}

func (s *service) DeletePost(postId, userId uint64) error {
	return s.repo.Write.DeletePost(postId, userId)
}

func (s *service) DeleteUserPosts(userId uint64) error {
	return s.repo.Write.DeleteUserPosts(userId)
}

func (s *service) GetPost(postIds []uint64) ([]*pb.Post, error) {
	var (
		posts   []*pb.Post
		users   map[uint64]*pb.User
		userIds []uint64

		err error
	)

	{
		if len(postIds) == 0 {
			return []*pb.Post{}, nil
		}
	}
	{
		rawPosts, err := s.repo.Read.GetPost(postIds)
		if err != nil {
			return nil, err
		}
		posts = make([]*pb.Post, len(rawPosts))
		for idx, p := range rawPosts {
			posts[idx] = &pb.Post{
				Id:      p.ID,
				Title:   p.Title,
				Body:    p.Body,
				OwnerId: p.OwnerId,
			}
		}
	}
	{
		notUniqueIds := make([]uint64, len(posts))
		for idx, i := range posts {
			notUniqueIds[idx] = i.OwnerId
		}
		userIds = utils.UniqueIds(notUniqueIds)
	}
	{
		users, err = s.publishers.GetUsers(userIds)
		if err != nil {
			return []*pb.Post{}, err
		}
	}
	{
		for _, p := range posts {
			p.User = users[p.OwnerId]
		}
	}

	return posts, err
}
