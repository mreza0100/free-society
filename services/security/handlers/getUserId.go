package handlers

import (
	"context"
	pb "freeSociety/proto/generated/security"
)

func (h *handlers) GetUserId(_ context.Context, in *pb.GetUserIdRequest) (*pb.GetUserIdResponse, error) {
	userId, err := h.srv.GetUserId(in.Token)

	return &pb.GetUserIdResponse{
		UserId: userId,
	}, err
}
