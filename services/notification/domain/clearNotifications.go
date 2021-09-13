package domain

func (s *service) ClearNotifications(userId uint64) error {
	cc, err := s.repo.Write.ClearNotifications(userId)
	cc.Commit()
	return err
}
