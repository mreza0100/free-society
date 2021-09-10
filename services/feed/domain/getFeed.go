package domain

func (s *service) GetFeed(userId, offset, limit uint64) ([]string, error) {
	return s.repo.Read.GetFeed(userId, offset, limit)
}
