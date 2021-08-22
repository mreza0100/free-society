package domain

func (s *service) IsPostsExists(postIds []uint64) ([]uint64, error) {
	return s.repo.Read.IsExists(postIds)
}
