package domain

import "freeSociety/services/notification/models"

func (s *service) GetNotifications(userId uint64, limit, offset int64) ([]models.Notification, error) {
	if limit > 50 {
		limit = 50
	}
	return s.repo.Read.GetNotifications(userId, limit, offset)
}
