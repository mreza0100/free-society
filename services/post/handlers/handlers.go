package handlers

import (
	"context"
	pb "microServiceBoilerplate/proto/generated/post"
	"microServiceBoilerplate/services/post/types"

	"github.com/mreza0100/golog"
)

type HandlersOptns struct {
	Srv        types.Sevice
	Lgr        *golog.Core
	Publishers types.Publishers
}

func NewHandlers(opts *HandlersOptns) types.Handlers {
	return &handlers{
		srv:        opts.Srv,
		lgr:        opts.Lgr,
		publishers: opts.Publishers,
	}
}

type handlers struct {
	srv        types.Sevice
	lgr        *golog.Core
	publishers types.Publishers

	pb.UnimplementedPostServiceServer
}

func (h *handlers) NewPost(_ context.Context, in *pb.NewPostRequest) (*pb.NewPostResponse, error) {
	postId, err := h.srv.NewPost(in.Title, in.Body, in.UserId)

	h.publishers.NewPost(in.UserId, postId)

	return &pb.NewPostResponse{Id: postId}, err

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
