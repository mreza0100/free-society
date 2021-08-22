package instances

import "freeSociety/services/post/models"

type read interface {
	GetPost(postIds []uint64) ([]*models.Post, error)
	IsExists(postIds []uint64) ([]uint64, error)
	IsPictureExist(name string) (bool, error)
}

type write interface {
	NewPost(title, body string, userId uint64, picturesName []string) (uint64, error)
	DeletePost(postId, userId uint64) (picturesName string, err error)
	DeleteUserPosts(userId uint64) ([]struct{ PicturesName string }, error)
}

type Repository struct {
	Read  read
	Write write
}
