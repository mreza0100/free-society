package domain

func (s *service) SetPost(userId, postId uint64) error {
	followers, err := s.publishers.GetFollowers(userId)
	if err != nil {
		return err
	}

	return s.repo.Write.SetPostOnFeeds(userId, postId, followers)
}