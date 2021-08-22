package handlers

import (
	"context"
	pb "freeSociety/proto/generated/post"
)

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
