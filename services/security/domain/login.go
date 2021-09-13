package domain

import (
	"errors"
	"freeSociety/configs"
	"freeSociety/services/security/utils"
	generalUtils "freeSociety/utils"
	dbhelper "freeSociety/utils/dbHelper"
	"freeSociety/utils/security"
	"time"
)

func (s *service) Login(userId uint64, device, password string) (token string, err error) {
	var (
		cc1 dbhelper.CommandController
	)
	defer func() {
		cc1.Done(err)
	}()

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

		_, cc, err := s.postgresRepo.Write.NewSession(userId, device, token, expire)
		cc1 = cc
		if err != nil {
			cc.Rollback()
			return "", err
		}
	}
	{
		err := s.redisRepo.Write.NewSession(token, userId)
		if err != nil {
			cc1.Rollback()
			return "", err
		}
	}

	return token, cc1.Done(nil)
}
