package domain

import (
	"freeSociety/configs"
	"freeSociety/utils/files/costume"
	"strings"
)

func (s *service) DeletePost(postId, userId uint64) error {
	rawPicturesNames, err := s.repo.Write.DeletePost(postId, userId)
	if err != nil {
		return err
	}

	picturesNames := strings.Split(rawPicturesNames, configs.DB_picture_sep)
	for i := 0; i < len(picturesNames); i++ {
		if err = costume.DeletPicture(picturesNames[i]); err != nil {
			return err
		}
	}

	return nil
}
