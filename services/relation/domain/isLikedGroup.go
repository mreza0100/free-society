package domain

func (s *service) IsLikedGroup(likerId uint64, postIds []uint64) ([]uint64, error) {
	return s.repo.Likes_read.IsLikedGroup(likerId, postIds)
}
