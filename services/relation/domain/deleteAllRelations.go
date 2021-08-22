package domain

func (s *service) DeleteAllRelations(userId uint64) error {
	return s.repo.Followers_write.DeleteAllRelations(userId)
}
