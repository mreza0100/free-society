package domain

func (s *service) GetUserId(token string) (uint64, error) {
	return s.redisRepo.Read.GetSession(token)
}
