package repository

import (
	dbhelper "freeSociety/utils/dbHelper"

	"github.com/mreza0100/golog"
	"gorm.io/gorm"
)

type write struct {
	lgr *golog.Core
	db  *gorm.DB
}

func (w *write) SetLikeNotification(userId, likerId uint64, postId string) (notificationId uint64, cc dbhelper.CommandController, err error) {
	cc, err = dbhelper.Transaction(w.db, func(tx *gorm.DB) error {
		const query = `INSERT INTO notifications (user_id, is_like, liker_id, post_id) VALUES (?, ?, ?, ?) RETURNING id`
		params := []interface{}{userId, true, likerId, postId}

		tx = w.db.Raw(query, params...)
		if tx.Error != nil {
			w.lgr.ErrorLog("Failed to set like notification", "error", tx.Error)
		}
		result := new(struct{ Id uint64 })
		tx.Scan(result)
		notificationId = likerId

		return tx.Error
	})

	return notificationId, cc, err
}

func (w *write) ClearNotifications(userId uint64) (cc dbhelper.CommandController, err error) {
	cc, err = dbhelper.Transaction(w.db, func(tx *gorm.DB) error {
		const query = `UPDATE notifications SET seen = true WHERE user_id = ?`
		params := []interface{}{userId}

		tx = w.db.Exec(query, params...)
		return tx.Error
	})

	return cc, err
}
