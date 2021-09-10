package models

type Post struct {
	ID      string `bson:"_id"`
	OwnerId uint64

	Title string
	Body  string

	PicturesName []string

	CreatedAt string
}
