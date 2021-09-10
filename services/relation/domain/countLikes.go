package domain

import "freeSociety/services/relation/instances"

func (s *service) CountLikes(postIds []string) (instances.CountResult, error) {
	return s.repo.Likes_read.CountLikes(postIds)
}
