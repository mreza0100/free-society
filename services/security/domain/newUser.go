package domain

import (
	"freeSociety/configs"
	"freeSociety/services/security/utils"
	generalUtils "freeSociety/utils"
	dbhelper "freeSociety/utils/dbHelper"
	"freeSociety/utils/security"
	"time"
)

func (s *service) NewUser(userId uint64, device, password string) (token string, err error) {
	var (
		cc1 dbhelper.CommandController
		cc2 dbhelper.CommandController
	)
	defer func() {
		cc1.Done(err)
		cc2.Done(err)
	}()

	{
		hashPass := security.HashSha1(password)
		var err error
		cc1, err = s.postgresRepo.Write.NewUser(userId, hashPass)
		if err != nil {
			return "", err
		}
	}
	{
		token = utils.CreateToken()
		expire := generalUtils.ParseDateForDb(time.Now().Add(configs.Token_expire))

		var err error
		_, cc2, err = s.postgresRepo.Write.NewSession(userId, device, token, expire)
		if err != nil {
			return "", err
		}
	}
	{
		err := s.redisRepo.Write.NewSession(token, userId)
		if err != nil {
			return "", err
		}
	}

	return token, nil
}
