package domain

func (s *service) DeleteAllRelations(userId uint64) error {
	cc, err := s.repo.Followers_write.DeleteAllRelations(userId)
	cc.Commit()
	return err
}
