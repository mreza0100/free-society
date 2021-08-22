package domain

func (s *service) IsFollowing(follower uint64, followings []uint64) (map[uint64]bool, error) {
	var (
		result   map[uint64]interface{}
		response map[uint64]bool

		err error
	)

	{
		result, err = s.repo.Followers_read.IsFollowingGroup(follower, followings)
		if err != nil {
			return nil, err
		}
	}
	{
		response = make(map[uint64]bool, len(followings))

		for _, i := range followings {
			_, isFollowing := result[i]
			response[i] = isFollowing
		}
	}

	return response, nil
}
