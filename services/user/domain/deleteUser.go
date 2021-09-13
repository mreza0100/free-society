package domain

import (
	"freeSociety/configs"
	dbhelper "freeSociety/utils/dbHelper"
	"freeSociety/utils/files/costume"
)

func (s *service) DeleteUser(id uint64) (fErr error) {
	var (
		avatarName string
		cc         dbhelper.CommandController
	)
	defer func() {
		cc.Done(fErr)
	}()

	{
		errsCh := make(chan error, 2)

		// flow request to services.DeleteUser and publish request to post service to delete user posts
		go func() { var err error; avatarName, cc, err = s.repo.Write.DeleteUser(id); errsCh <- err }()
		go func() { errsCh <- s.publishers.DeleteUser(id) }()

		for i := 0; i < cap(errsCh); i++ {
			if err := <-errsCh; err != nil {
				return err
			}
		}
	}

	if avatarName == configs.MaleDefaultAvatarPath || avatarName == configs.FemaleDefaultAvatarPath {
		return nil
	}

	return costume.DeleteAvatar(avatarName)
}
