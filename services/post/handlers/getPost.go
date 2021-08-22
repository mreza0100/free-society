package handlers

import (
	"context"
	pb "freeSociety/proto/generated/post"
)

func (h *handlers) GetPost(_ context.Context, in *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	posts, err := h.srv.GetPost(in.RequestorId, in.Ids)

	return &pb.GetPostResponse{Posts: posts}, err
}
