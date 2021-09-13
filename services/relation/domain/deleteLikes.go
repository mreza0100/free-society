package domain

func (s *service) DeleteLikes(liker uint64) error {
	cc, err := s.repo.Likes_write.PurgeUserLikes(liker)
	cc.Commit()
	return err
}
