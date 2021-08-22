package domain

func (s *service) DeleteLikes(liker uint64) error {
	return s.repo.Likes_write.PurgeUserLikes(liker)
}
