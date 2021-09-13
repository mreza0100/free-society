package domain

func (s *service) Logout(token string) (err error) {
	errRedis := s.redisRepo.Write.DeleteSession(token)
	postgresCc, errPostgres := s.postgresRepo.Write.DeleteSessionByToken(token)

	defer func() {
		postgresCc.Done(err)
	}()

	if errRedis != nil {
		return errRedis
	}

	return errPostgres
}
