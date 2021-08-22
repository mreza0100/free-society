package domain

import (
	"freeSociety/services/user/models"
	"freeSociety/utils/files/costume"
)

func (s *service) GetUsers(ids []uint64) (map[uint64]*models.User, error) {
	var (
		rawUsers []*models.User
		result   map[uint64]*models.User
		err      error
	)

	{
		rawUsers, err = s.repo.Read.GetUsersByIds(ids)
		if err != nil {
			return nil, err
		}
	}

	{
		result = make(map[uint64]*models.User, len(rawUsers))
		for _, u := range rawUsers {
			result[u.ID] = &models.User{
				Name:       u.Name,
				Email:      u.Email,
				ID:         u.ID,
				Gender:     u.Gender,
				AvatarName: costume.ExportAvatar(u.AvatarName),
				CreatedAt:  u.CreatedAt,
			}
		}
	}

	return result, err
}
