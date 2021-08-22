package domain

import (
	"freeSociety/configs"
	"freeSociety/utils/files/costume"
	"strings"
)

func (s *service) DeleteUserPosts(userId uint64) error {
	picturesName, err := s.repo.Write.DeleteUserPosts(userId)
	if err != nil {
		return err
	}

	for _, rawPicNmaes := range picturesName {
		for _, picName := range strings.Split(rawPicNmaes.PicturesName, configs.DB_picture_sep) {
			if costume.DeletPicture(picName) != nil {
				return err
			}
		}
	}

	return nil
}
