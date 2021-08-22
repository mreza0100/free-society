package domain

func (s *service) PurgeUser(userId uint64) error {
	var (
		tokens []string
		chErr  = make(chan error)
	)

	{
		sessions, err := s.postgresRepo.Write.DeleteUserSessions(userId)
		if err != nil {
			return err
		}
		tokens = make([]string, len(sessions))
		for idx, i := range sessions {
			tokens[idx] = i.Token
		}
	}
	{
		go func(ch chan error) {
			ch <- s.redisRepo.Write.DeleteSession(tokens...)
		}(chErr)
	}
	{
		go func(ch chan error) {
			ch <- s.postgresRepo.Write.DeletePassword(userId)
		}(chErr)
	}

	for i := 0; i < 2; i++ {
		if err := <-chErr; err != nil {
			return err
		}
	}

	return nil
}
