package instances

import "freeSociety/services/notification/models"

type read interface {
	GetNotifications(userId uint64, limit, offset int64) ([]models.Notification, error)
}

type write interface {
	SetLikeNotification(userId, likerId, postId uint64) (notificationId uint64, err error)
	ClearNotifications(userId uint64) error
}

type Repository struct {
	Read  read
	Write write
}
