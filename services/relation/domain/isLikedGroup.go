package domain

func (s *service) IsLikedGroup(likerId uint64, postIds []string) ([]string, error) {
	return s.repo.Likes_read.IsLikedGroup(likerId, postIds)
}
