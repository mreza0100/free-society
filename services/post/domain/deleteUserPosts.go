package domain

import (
	"freeSociety/utils/files/costume"
)

func (s *service) DeleteUserPosts(userId uint64) error {
	picturesName, err := s.repo.Write.DeleteUserPosts(userId)
	if err != nil {
		return err
	}

	for _, rawPicNmaes := range picturesName {
		for _, picName := range rawPicNmaes.PicturesName {
			if costume.DeletPicture(picName) != nil {
				return err
			}
		}
	}

	return nil
}
