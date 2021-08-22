package handlers

import (
	"context"
	pb "freeSociety/proto/generated/notification"
)

func (h *handlers) ClearNotifications(_ context.Context, in *pb.ClearNotificationsRequest) (*pb.ClearNotificationsResponse, error) {
	err := h.srv.ClearNotifications(in.UserId)

	return &pb.ClearNotificationsResponse{}, err

}
