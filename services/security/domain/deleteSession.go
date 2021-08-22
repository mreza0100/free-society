package domain

func (s *service) DeleteSession(sessionId uint64) (err error) {
	var (
		token string
	)

	{
		session, err := s.postgresRepo.Write.DeleteSessionById(sessionId)
		if err != nil {
			return err
		}
		token = session.Token
	}
	{
		err = s.redisRepo.Write.DeleteSession(token)
		if err != nil {
			return err
		}
	}
	return nil
}
