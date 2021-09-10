package domain

func (s *service) SetLikeNotification(userId, likerId uint64, postId string) (uint64, error) {
	return s.repo.Write.SetLikeNotification(userId, likerId, postId)
}
