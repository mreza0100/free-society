package handlers

import (
	"context"
	pb "freeSociety/proto/generated/security"
)

func (h *handlers) GetUserId(_ context.Context, in *pb.GetUserIdRequest) (*pb.GetUserIdResponse, error) {
	var (
		userId uint64
		err    error
		result *pb.GetUserIdResponse
	)

	{
		userId, err = h.srv.GetUserId(in.Token)
	}

	{
		result = &pb.GetUserIdResponse{
			UserId: userId,
		}
	}

	return result, err
}
