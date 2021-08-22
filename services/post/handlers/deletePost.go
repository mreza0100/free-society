package handlers

import (
	"context"
	pb "freeSociety/proto/generated/post"
)

func (h *handlers) DeletePost(_ context.Context, in *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	return &pb.DeletePostResponse{}, h.srv.DeletePost(in.PostId, in.UserId)
}
