package domain

import (
	"errors"
	dbhelper "freeSociety/utils/dbHelper"
	"freeSociety/utils/security"
)

func (s *service) ChangePassword(userId uint64, prevPassword, newPassword string) (err error) {
	var (
		tokens []string
		cc1    dbhelper.CommandController
		cc2    dbhelper.CommandController
	)
	defer func() {
		cc1.Done(err)
		cc2.Done(err)
	}()

	{
		hashPass, err := s.postgresRepo.Read.GetHashPass(userId)
		if err != nil {
			return err
		}

		if !security.HashSha1Compare(hashPass, prevPassword) {
			return errors.New("password is wrong")
		}
	}

	{
		cc, err := s.postgresRepo.Write.ChangeHashPass(userId, security.HashSha1(newPassword))
		cc1 = cc
		if err != nil {
			return err
		}
	}

	{
		sessions, cc, err := s.postgresRepo.Write.DeleteUserSessions(userId)
		cc2 = cc
		if err != nil {
			return err
		}

		tokens = make([]string, len(sessions))
		for idx, i := range sessions {
			tokens[idx] = i.Token
		}
	}

	return s.redisRepo.Write.DeleteSession(tokens...)
}
