package domain

import (
	"freeSociety/configs"
	"freeSociety/utils/files/costume"
)

func (s *service) DeleteUser(id uint64) error {
	avatarName, err := s.repo.Write.DeleteUser(id)
	if err != nil {
		return err
	}

	if avatarName == configs.MaleDefaultAvatarPath || avatarName == configs.FemaleDefaultAvatarPath {
		return nil
	}

	return costume.DeletAvatar(avatarName)
}
