package domain

func (s *service) DeleteFeed(userId uint64) error {
	return s.repo.Write.DeleteFeed(userId)
}
