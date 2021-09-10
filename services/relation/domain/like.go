package domain

func (s *service) Like(likerId, ownerId uint64, postId string) error {
	return s.repo.Likes_write.Like(likerId, ownerId, postId)
}
