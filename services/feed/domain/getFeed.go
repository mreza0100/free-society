package domain

func (s *service) GetFeed(userId, offset, limit uint64) ([]uint64, error) {
	return s.repo.Read.GetFeed(userId, offset, limit)
}
