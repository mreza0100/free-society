package instances

import (
	"freeSociety/services/notification/models"
	dbhelper "freeSociety/utils/dbHelper"
)

type read interface {
	GetNotifications(userId uint64, limit, offset int64) ([]models.Notification, error)
}

type write interface {
	SetLikeNotification(userId, likerId uint64, postId string) (notificationId uint64, cc dbhelper.CommandController, err error)
	ClearNotifications(userId uint64) (cc dbhelper.CommandController, err error)
}

type Repository struct {
	Read  read
	Write write
}
