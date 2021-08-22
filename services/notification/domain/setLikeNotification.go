package domain

func (s *service) SetLikeNotification(userId, likerId, postId uint64) (uint64, error) {
	return s.repo.Write.SetLikeNotification(userId, likerId, postId)
}
