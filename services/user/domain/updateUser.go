package domain

import (
	"freeSociety/configs"
	"freeSociety/utils"
	dbhelper "freeSociety/utils/dbHelper"
	"freeSociety/utils/files/costume"
)

func (s *service) UpdateUser(userId uint64, name, gender, avatarFormat string, avatar []byte) (err error) {
	var (
		avatarName string
		cc         dbhelper.CommandController
	)

	defer func() {
		cc.Done(err)
	}()

	{
		prevData, err := s.repo.Read.GetUserById(userId)
		if err != nil {
			return err
		}
		if !(prevData.AvatarName == configs.MaleDefaultAvatarPath || prevData.AvatarName == configs.FemaleDefaultAvatarPath) {
			err = costume.DeleteAvatar(prevData.AvatarName)
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

	cc, err = s.repo.Write.UpdateUser(userId, name, gender, avatarName)
	if err != nil {
		return err
	}
	return costume.SaveAvatar(avatarName, avatar)
}
