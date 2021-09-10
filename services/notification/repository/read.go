package repository

import (
	"freeSociety/services/notification/models"

	"github.com/mreza0100/golog"
	"gorm.io/gorm"
)

type read struct {
	lgr *golog.Core
	db  *gorm.DB
}

func (r *read) GetNotifications(userId uint64, limit, offset int64) ([]models.Notification, error) {
	const query = "SELECT * FROM notifications WHERE user_id = ? limit ? offset ?"
	params := []interface{}{userId, limit, offset}

	tx := r.db.Raw(query, params...)

	if tx.Error != nil {
		return nil, tx.Error
	}

	notifications := make([]models.Notification, 0)
	tx.Scan(&notifications)

	return notifications, nil
}
