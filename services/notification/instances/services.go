package instances

import "freeSociety/services/notification/models"

type Sevice interface {
	SetLikeNotification(userId, likerId, postId uint64) (notificationId uint64, err error)
	Getnotifications(userId uint64, limit, offset int) ([]models.Notification, error)
	ClearNotifications(userId uint64) error
}
