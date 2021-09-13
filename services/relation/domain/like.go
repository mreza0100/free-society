package domain

func (s *service) Like(likerId, ownerId uint64, postId string) error {
	cc, err := s.repo.Likes_write.Like(likerId, ownerId, postId)
	cc.Commit()
	return err
}
