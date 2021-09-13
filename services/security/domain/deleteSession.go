package domain

import dbhelper "freeSociety/utils/dbHelper"

func (s *service) DeleteSession(sessionId uint64) (err error) {
	var (
		token string
		cc1   dbhelper.CommandController
	)
	defer func() {
		cc1.Done(err)
	}()

	{
		session, cc, err := s.postgresRepo.Write.DeleteSessionById(sessionId)
		cc1 = cc
		if err != nil {
			return err
		}
		token = session.Token
	}
	{
		err = s.redisRepo.Write.DeleteSession(token)
		if err != nil {
			return cc1.Done(err)
		}
	}

	return cc1.Done(nil)
}
