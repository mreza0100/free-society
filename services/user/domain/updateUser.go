package domain

import (
	"freeSociety/configs"
	"freeSociety/utils"
	"freeSociety/utils/files/costume"
)

func (s *service) UpdateUser(userId uint64, name, gender, avatarFormat string, avatar []byte) error {
	var (
		avatarName string
	)

	{
		prevData, err := s.repo.Read.GetUserById(userId)
		if err != nil {
			return err
		}
		if !(prevData.AvatarName == configs.MaleDefaultAvatarPath || prevData.AvatarName == configs.FemaleDefaultAvatarPath) {
			err = costume.DeletAvatar(prevData.AvatarName)
			if err != nil {
				return err
			}
		}
	}
	{
		if len(avatar) != 0 {
			id := utils.GenerateUuid()
			avatarName = id + avatarFormat
		} else {
			if gender == "male" {
				avatarName = configs.MaleDefaultAvatarPath
			} else {
				avatarName = configs.FemaleDefaultAvatarPath
			}
		}
	}

	err := s.repo.Write.UpdateUser(userId, name, gender, avatarName)
	if err != nil {
		return err
	}
	return costume.SaveAvatar(avatarName, avatar)

}
