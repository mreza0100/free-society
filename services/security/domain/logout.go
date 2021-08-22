package domain

func (s *service) Logout(token string) error {
	// must delete from both
	errRedis := s.redisRepo.Write.DeleteSession(token)
	errPostgres := s.postgresRepo.Write.DeleteSessionByToken(token)

	if errRedis != nil {
		return errRedis
	}

	return errPostgres
}
