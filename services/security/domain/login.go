package domain

import (
	"errors"
	"freeSociety/configs"
	"freeSociety/services/security/utils"
	generalUtils "freeSociety/utils"
	"freeSociety/utils/security"
	"time"
)

func (s *service) Login(userId uint64, device, password string) (string, error) {
	var (
		token string
		err   error
	)

	{
		hashPass, err := s.postgresRepo.Read.GetHashPass(userId)
		if err != nil {
			return "", errors.New("email or password is wrong")
		}
		if !security.HashSha1Compare(hashPass, password) {
			return "", errors.New("email or password is wrong")
		}
	}
	{
		token = utils.CreateToken()
	}
	{
		expire := generalUtils.ParseDateForDb(time.Now().Add(configs.Token_expire))

		_, err = s.postgresRepo.Write.NewSession(userId, device, token, expire)
		if err != nil {
			return "", err
		}
	}
	{
		err = s.redisRepo.Write.NewSession(token, userId)
		if err != nil {
			return "", err
		}
	}

	return token, nil
}
