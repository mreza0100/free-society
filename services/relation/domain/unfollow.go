package domain

func (s *service) Unfollow(following, follower uint64) error {
	return s.repo.Followers_write.RemoveFollow(following, follower)
}
