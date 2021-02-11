package handlers

import (
	"context"
	pb "microServiceBoilerplate/proto/generated/post"
	"microServiceBoilerplate/services/post/domain"

	"github.com/mreza0100/golog"
)

type Handlers interface {
	pb.PostServiceServer

	DeleteUserPosts(userId uint64) error
}

func NewHandlers(srv domain.Sevice, Lgr *golog.Core) Handlers {
	return &handlers{
		srv: srv,
		lgr: Lgr,
	}
}

type handlers struct {
	srv domain.Sevice
	lgr *golog.Core

	pb.UnimplementedPostServiceServer
}

func (h *handlers) NewPost(_ context.Context, in *pb.NewPostRequest) (*pb.NewPostResponse, error) {
	id, err := h.srv.NewPost(in.Title, in.Body, in.UserId)

	return &pb.NewPostResponse{Id: id}, err

}

func (h *handlers) DeletePost(_ context.Context, in *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	return &pb.DeletePostResponse{}, h.srv.DeletePost(in.PostId, in.UserId)
}

func (h *handlers) GetPost(_ context.Context, in *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	res, err := h.srv.GetPost(in.Ids)

	if err != nil {
		return &pb.GetPostResponse{}, err
	}

	return &pb.GetPostResponse{Posts: res}, nil

}
func (h *handlers) DeleteUserPosts(userId uint64) error {
	return h.srv.DeleteUserPosts(userId)
}
