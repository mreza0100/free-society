package domain

import (
	"freeSociety/configs"
	"freeSociety/utils"
	"freeSociety/utils/files/costume"
)

func (s *service) NewUser(name, email, gender, avatarFormat string, avatar []byte) (uint64, error) {
	var (
		avatarName string
		isCostume  bool
	)

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

	userId, err := s.repo.Write.NewUser(name, gender, email, avatarName)
	if err != nil {
		return 0, err
	}

	if isCostume {
		err = costume.SaveAvatar(avatarName, avatar)
	}

	return userId, err
}
