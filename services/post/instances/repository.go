package instances

import "microServiceBoilerplate/services/post/models"

type read interface {
	GetPost(postIds []uint64) ([]*models.Post, error)
	IsExists(postIds []uint64) ([]uint64, error)
}

type write interface {
	NewPost(title, body string, userId uint64) (uint64, error)
	DeletePost(postId, userId uint64) error
	DeleteUserPosts(userId uint64) error
}

type Repository struct {
	Read  read
	Write write
}
