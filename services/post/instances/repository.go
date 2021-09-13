package instances

import "freeSociety/services/post/models"

type read interface {
	GetPost(postIds []string) ([]*models.Post, error)
	IsExists(postIds []string) ([]string, error)
	IsPictureExist(name string) (bool, error)
}

type write interface {
	NewPost(title, body string, userId uint64, picturesNames []string) (postId string, err error)
	DeletePost(postId string, userId uint64) (picturesNames []string, err error)
	DeleteUserPosts(userId uint64) (PicturesNames []struct{ PicturesName []string }, err error)
}

type Repository struct {
	Read  read
	Write write
}
