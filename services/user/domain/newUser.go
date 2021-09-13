package domain

import (
	"freeSociety/configs"
	"freeSociety/utils"
	dbhelper "freeSociety/utils/dbHelper"
	"freeSociety/utils/files/costume"
)

func (s *service) NewUser(name, email, gender, avatarFormat string, avatar []byte) (userId uint64, err error) {
	var (
		avatarName string
		isCostume  bool
		cc         dbhelper.CommandController
	)

	defer func() {
		cc.Done(err)
	}()

	{
		if len(avatar) == 0 {
			if gender == "male" {
				avatarName = configs.MaleDefaultAvatarPath
			} else {
				avatarName = configs.FemaleDefaultAvatarPath
			}
			isCostume = false
		} else {
			id := utils.GenerateUuid()
			avatarName = id + avatarFormat
			isCostume = true
		}
	}

	userId, cc, err = s.repo.Write.NewUser(name, gender, email, avatarName)
	if err != nil {
		return 0, err
	}

	if isCostume {
		err = costume.SaveAvatar(avatarName, avatar)
	}

	return userId, err
}
