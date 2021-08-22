package domain

import "freeSociety/services/relation/instances"

func (s *service) CountLikes(postIds []uint64) (instances.CountResult, error) {
	return s.repo.Likes_read.CountLikes(postIds)
}
