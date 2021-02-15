package domain

import (
	"errors"
	"microServiceBoilerplate/services/security/db"
	"microServiceBoilerplate/services/security/types"
	"microServiceBoilerplate/services/security/utils"
	"microServiceBoilerplate/utils/security"

	"github.com/mreza0100/golog"
)

type ServiceOpts struct {
	Lgr *golog.Core
}

func NewService(opts ServiceOpts) types.Sevice {
	psDAOS := &db.PS_DAOS{
		Lgr: opts.Lgr.With("in Postgres DAOS: "),
	}
	redisDAOS := &db.RedisDAOS{
		Lgr: opts.Lgr.With("in Redis DAOS: "),
	}

	return &service{
		PS_DAOS:   psDAOS,
		redisDAOS: redisDAOS,
		Lgr:       opts.Lgr.With("In domain: "),
	}
}

type service struct {
	redisDAOS *db.RedisDAOS
	PS_DAOS   *db.PS_DAOS
	Lgr       *golog.Core
}

func (s *service) NewUser(userId uint64, device, password string) (token string, err error) {
	debug := s.Lgr.DebugPKG("NewUser")

	{
		hashPass := security.HashIt(password)
		err = s.PS_DAOS.NewUser(userId, hashPass)
		debug("after s.PS_DAOS.NewUser")(err)
		if err != nil {
			return "", err
		}
	}
	{
		token = utils.CreateToken()
		_, err = s.PS_DAOS.NewSession(userId, device, token)
		debug("after s.PS_DAOS.NewSession")(err)
		if err != nil {
			return "", err
		}
	}
	{
		err = s.redisDAOS.NewSession(token, userId)
		debug("after s.redisDAOS.NewSession")(err)
		if err != nil {
			return "", err
		}
	}

	return token, nil
}

func (s *service) Login(userId uint64, device, password string) (string, error) {
	var (
		token string
		err   error
	)

	{
		hashPass, err := s.PS_DAOS.GetHashPass(userId)
		if err != nil {
			return "", errors.New("email or password is wrong")
		}
		if !security.HashCompare(hashPass, password) {
			return "", errors.New("email or password is wrong")
		}
	}
	{
		token = utils.CreateToken()
	}
	{
		_, err = s.PS_DAOS.NewSession(userId, device, token)
		if err != nil {
			return "", err
		}
	}
	{
		err = s.redisDAOS.NewSession(token, userId)
		if err != nil {
			return "", err
		}
	}

	return token, nil
}

func (s *service) Logout(token string) (err error) {
	{
		err = s.redisDAOS.DeleteSession(token)
	}
	{
		err = s.PS_DAOS.DeleteSessionByToken(token)
	}
	return err
}

func (s *service) GetUserId(token string) (uint64, error) {
	return s.redisDAOS.GetSession(token)
}
