package domain

func (s *service) UndoLike(likerId uint64, postId string) error {
	return s.repo.Likes_write.UndoLike(likerId, postId)
}
