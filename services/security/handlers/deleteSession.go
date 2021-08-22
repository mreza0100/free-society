package handlers

import (
	"context"
	pb "freeSociety/proto/generated/security"
)

func (h *handlers) DeleteSession(_ context.Context, in *pb.DeleteSessionRequest) (*pb.DeleteSessionResponse, error) {
	return &pb.DeleteSessionResponse{}, h.srv.DeleteSession(in.SessionId)
}
