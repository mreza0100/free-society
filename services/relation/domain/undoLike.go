package domain

func (s *service) UndoLike(likerId, postId uint64) error {
	return s.repo.Likes_write.UndoLike(likerId, postId)
}
