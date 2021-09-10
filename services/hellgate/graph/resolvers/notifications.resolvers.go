package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	pbNotification "freeSociety/proto/generated/notification"
	models "freeSociety/services/hellgate/graph/model"
	"freeSociety/services/hellgate/security"
)

func (r *mutationResolver) ClearNotifications(ctx context.Context) (bool, error) {
	userId := security.GetUserId(ctx)

	_, err := r.notificationConn.ClearNotifications(ctx, &pbNotification.ClearNotificationsRequest{
		UserId: userId,
	})

	return err == nil, err
}

func (r *queryResolver) GetNotifications(ctx context.Context, offset int, limit int) ([]*models.Notification, error) {
	userId := security.GetUserId(ctx)

	response, err := r.notificationConn.GetNotifications(ctx, &pbNotification.GetNotificationsRequest{
		UserId: userId,
		Offset: int64(offset),
		Limit:  int64(limit),
	})
	if err != nil {
		return nil, err
	}

	notifications := make([]*models.Notification, len(response.Notifications))
	for i := 0; i < len(response.Notifications); i++ {
		notifications[i] = &models.Notification{
			ID: int(response.Notifications[i].Id),

			IsLike:  response.Notifications[i].IsLike,
			LikerID: int(response.Notifications[i].LikerId),
			PostID:  response.Notifications[i].PostId,

			Seen: response.Notifications[i].Seen,
			Time: response.Notifications[i].Time,
		}
		r.Lgr.InfoLog(notifications[i].PostID)
	}

	return notifications, nil
}
