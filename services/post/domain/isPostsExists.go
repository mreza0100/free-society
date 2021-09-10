package domain

func (s *service) IsPostsExists(postIds []string) ([]string, error) {
	return s.repo.Read.IsExists(postIds)
}
