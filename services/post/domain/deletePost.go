package domain

import (
	"freeSociety/utils/files/costume"
)

func (s *service) DeletePost(postId string, userId uint64) error {
	rawPicturesNames, err := s.repo.Write.DeletePost(postId, userId)
	if err != nil {
		return err
	}

	for i := 0; i < len(rawPicturesNames); i++ {
		if err = costume.DeletPicture(rawPicturesNames[i]); err != nil {
			return err
		}
	}

	return nil
}
