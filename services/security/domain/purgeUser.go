package domain

import (
	dbhelper "freeSociety/utils/dbHelper"
)

func (s *service) PurgeUser(userId uint64) (err error) {
	var (
		tokens []string
		chErr  = make(chan error)
		cc1    dbhelper.CommandController
		cc2    dbhelper.CommandController
	)

	defer func() {
		cc1.Done(err)
		cc2.Done(err)
	}()

	{
		sessions, cc, err := s.postgresRepo.Write.DeleteUserSessions(userId)
		cc1 = cc
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
			cc, err := s.postgresRepo.Write.DeletePassword(userId)
			cc2 = cc
			ch <- err
		}(chErr)
	}

	for i := 0; i < 2; i++ {
		if err := <-chErr; err != nil {
			return err
		}
	}

	return nil
}
