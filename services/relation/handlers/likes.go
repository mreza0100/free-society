package handlers

import (
	"context"
	pb "microServiceBoilerplate/proto/generated/relation"
)

func (h *handlers) Like(_ context.Context, in *pb.LikeRequest) (*pb.LikeResponse, error) {
	return &pb.LikeResponse{}, h.srv.Like(in.LikerId, in.OwnerId, in.PostId)
}
