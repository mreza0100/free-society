package handlers

import (
	"context"
	pb "freeSociety/proto/generated/notification"
)

func (h *handlers) GetNotifications(_ context.Context, in *pb.GetNotificationsRequest) (*pb.GetNotificationsResponse, error) {
	rawNotifications, err := h.srv.GetNotifications(in.UserId, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}

	notifications := make([]*pb.Notification, len(rawNotifications))
	for i := 0; i < len(rawNotifications); i++ {
		notifications[i] = &pb.Notification{
			Id:      rawNotifications[i].ID,
			IsLike:  rawNotifications[i].IsLike,
			LikerId: rawNotifications[i].LikerId,
			PostId:  rawNotifications[i].PostId,
			Seen:    rawNotifications[i].Seen,
			Time:    rawNotifications[i].CreatedAt.String(),
		}
	}

	return &pb.GetNotificationsResponse{
		Notifications: notifications,
	}, nil
}
