package domain

func (s *service) GetFollowers(userId uint64) []uint64 {
	return s.repo.Followers_read.GetFollowers(userId)
}
