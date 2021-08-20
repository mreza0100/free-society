package handlers

import (
	"context"
	"freeSociety/configs"
	pb "freeSociety/proto/generated/post"
	"freeSociety/services/post/instances"
	"freeSociety/utils/files"

	"github.com/mreza0100/golog"
)

type NewOpts struct {
	Srv        instances.Sevice
	Lgr        *golog.Core
	Publishers instances.Publishers
}

func New(opts *NewOpts) instances.Handlers {
	if !files.FileExist(configs.PicturesPath) {
		files.CreateDir(configs.PicturesPath)
	}

	return &handlers{
		srv:        opts.Srv,
		lgr:        opts.Lgr,
		publishers: opts.Publishers,
	}
}

type handlers struct {
	srv        instances.Sevice
	lgr        *golog.Core
	publishers instances.Publishers

	pb.UnimplementedPostServiceServer
}

func (h *handlers) NewPost(_ context.Context, in *pb.NewPostRequest) (*pb.NewPostResponse, error) {
	postId, err := h.srv.NewPost(in.Title, in.Body, in.UserId, in.Pictures)
	if err != nil {
		return &pb.NewPostResponse{}, err
	}

	err = h.publishers.NewPost(in.UserId, postId)
	if err != nil {
		return &pb.NewPostResponse{}, err
	}

	return &pb.NewPostResponse{Id: postId}, nil
}

func (h *handlers) DeletePost(_ context.Context, in *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	return &pb.DeletePostResponse{}, h.srv.DeletePost(in.PostId, in.UserId)
}

func (h *handlers) GetPost(_ context.Context, in *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	posts, err := h.srv.GetPost(in.RequestorId, in.Ids)

	return &pb.GetPostResponse{Posts: posts}, err
}
