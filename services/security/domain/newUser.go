package domain

import (
	"freeSociety/configs"
	"freeSociety/services/security/utils"
	generalUtils "freeSociety/utils"
	"freeSociety/utils/security"
	"time"
)

func (s *service) NewUser(userId uint64, device, password string) (token string, err error) {
	debug, success := s.lgr.DebugPKG("NewUser", false)

	{
		hashPass := security.HashSha1(password)
		err = s.postgresRepo.Write.NewUser(userId, hashPass)
		if debug("after s.postgresRepo.NewUser")(err) != nil {
			return "", err
		}
	}
	{
		token = utils.CreateToken()
		expire := generalUtils.ParseDateForDb(time.Now().Add(configs.Token_expire))

		_, err = s.postgresRepo.Write.NewSession(userId, device, token, expire)
		if debug("after s.postgresRepo..NewSession")(err) != nil {
			return "", err
		}
	}
	{
		err = s.redisRepo.Write.NewSession(token, userId)
		if debug("after s.redisDAOS.NewSession")(err) != nil {
			return "", err
		}
	}

	success(token)
	return token, nil
}
