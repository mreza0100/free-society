package domain

import (
	"microServiceBoilerplate/services/relation/instances"
)

func (s *service) Like(likerId, ownerId, postId uint64) error {
	return s.repo.Likes_write.Like(likerId, ownerId, postId)
}

func (s *service) UndoLike(likerId, postId uint64) error {
	return s.repo.Likes_write.UndoLike(likerId, postId)
}

func (s *service) IsLikedGroup(likerId uint64, postIds []uint64) ([]uint64, error) {
	return s.repo.Likes_read.IsLikedGroup(likerId, postIds)
}

func (s *service) CountLikes(postIds []uint64) (instances.CountResult, error) {
	return s.repo.Likes_read.CountLikes(postIds)
}

func (s *service) DeleteLikes(liker uint64) error {
	return s.repo.Likes_write.PurgeUserLikes(liker)
}
