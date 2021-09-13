package domain

func (s *service) SetLikeNotification(userId, likerId uint64, postId string) (uint64, error) {
	id, cc, err := s.repo.Write.SetLikeNotification(userId, likerId, postId)
	cc.Commit()
	return id, err
}
