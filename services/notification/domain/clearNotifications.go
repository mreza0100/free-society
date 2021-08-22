package domain

func (s *service) ClearNotifications(userId uint64) error {
	return s.repo.Write.ClearNotifications(userId)
}
