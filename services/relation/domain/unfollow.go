package domain

func (s *service) Unfollow(following, follower uint64) error {
	cc, err := s.repo.Followers_write.RemoveFollow(following, follower)
	cc.Commit()
	return err
}
