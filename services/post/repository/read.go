package repository

import (
	"context"
	"freeSociety/services/post/models"

	"github.com/mreza0100/golog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type read struct {
	lgr         *golog.Core
	db          *mongo.Client
	pCollection *mongo.Collection
	write       *write
}

func (r *read) GetPost(rawPostIds []string) ([]*models.Post, error) {
	postIds := make([]primitive.ObjectID, len(rawPostIds))

	for i := 0; i < len(rawPostIds); i++ {
		var err error
		postIds[i], err = primitive.ObjectIDFromHex(rawPostIds[i])
		if err != nil {
			return nil, err
		}
	}

	query := bson.M{
		"_id": bson.M{"$in": postIds},
	}

	cursor, err := r.pCollection.Find(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	posts := make([]*models.Post, 0, len(postIds))
	if err := cursor.All(context.Background(), &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *read) IsExists(postIds []string) ([]string, error) {
	result, err := r.GetPost(postIds)
	if err != nil {
		return nil, err
	}

	exists := make([]string, 0, len(postIds))
	for _, post := range result {
		exists = append(exists, post.ID)
	}

	return exists, nil

}

func (r *read) IsPictureExist(name string) (bool, error) {
	query := bson.M{
		"PicturesName": name,
	}

	exists, err := r.pCollection.CountDocuments(context.Background(), query, options.Count().SetLimit(1))

	return exists > 0, err
}
