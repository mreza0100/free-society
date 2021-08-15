package handlers

import (
	"context"
	pb "freeSociety/proto/generated/notification"
	"freeSociety/services/notification/instances"

	"github.com/mreza0100/golog"
)

type NewOpts struct {
	Srv        instances.Sevice
	Lgr        *golog.Core
	Publishers instances.Publishers
}

func New(opts *NewOpts) instances.Handlers {
	return &handlers{
		srv:        opts.Srv,
		lgr:        opts.Lgr,
		publishers: opts.Publishers,
	}
}

type handlers struct {
	srv        instances.Sevice
	lgr        *golog.Core
	publishers instances.Publishers

	pb.UnimplementedNotificationServiceServer
}

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
func (h *handlers) ClearNotifications(_ context.Context, in *pb.ClearNotificationsRequest) (*pb.ClearNotificationsResponse, error) {
	err := h.srv.ClearNotifications(in.UserId)

	return &pb.ClearNotificationsResponse{}, err

}
