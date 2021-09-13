package domain

func (s *service) UndoLike(likerId uint64, postId string) error {
	cc, err := s.repo.Likes_write.UndoLike(likerId, postId)
	cc.Commit()
	return err
}
