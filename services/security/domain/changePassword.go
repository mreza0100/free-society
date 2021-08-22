package domain

import (
	"errors"
	"freeSociety/utils/security"
)

func (s *service) ChangePassword(userId uint64, prevPassword, newPassword string) error {
	var (
		tokens []string
	)

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
		err := s.postgresRepo.Write.ChangeHashPass(userId, security.HashSha1(newPassword))
		if err != nil {
			return err
		}
	}

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

	return s.redisRepo.Write.DeleteSession(tokens...)
}
