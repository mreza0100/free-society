package handlers

import (
	"context"
	pb "freeSociety/proto/generated/security"
	"freeSociety/services/security/models"
)

func (h *handlers) GetSessions(_ context.Context, in *pb.GetSessionsRequest) (*pb.GetSessionsResponse, error) {
	var (
		result    []*models.Session
		converted []*pb.Session
		err       error
	)

	{
		result, err = h.srv.GetSessions(in.UserId)
		if err != nil {
			return &pb.GetSessionsResponse{}, err
		}
	}

	{
		converted = make([]*pb.Session, len(result))
		for idx, i := range result {
			converted[idx] = &pb.Session{
				SessionId: i.ID,
				Device:    i.Device,
				CreatedAt: uint64(i.CreatedAt.Unix()),
			}
		}
	}

	return &pb.GetSessionsResponse{
		Sessions: converted,
	}, err
}
