package handlers

import (
	"context"
	pb "freeSociety/proto/generated/user"
)

func (h *handlers) GetUser(_ context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return h.srv.GetUser(in.RequestorId, in.Id, in.Email)
}
