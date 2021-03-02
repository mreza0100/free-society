package domain

func (s *service) Like(likerId, ownerId, postId uint64) error {
	return s.repo.Likes_write.Like(likerId, ownerId, postId)
}

func (s *service) UndoLike(likerId, postId uint64) error {
	return s.repo.Likes_write.UndoLike(likerId, postId)
}
