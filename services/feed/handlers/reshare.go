package handlers

import (
	"context"
	pb "freeSociety/proto/generated/feed"
)

func (h *handlers) Reshare(_ context.Context, in *pb.ReshareRequest) (*pb.ReshareResponse, error) {
	return &pb.ReshareResponse{}, h.srv.Reshare(in.UserId, in.PostId)
}
