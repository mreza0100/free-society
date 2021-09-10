package repository

import (
	"context"
	"errors"
	"freeSociety/services/post/models"
	"time"

	"github.com/mreza0100/golog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type write struct {
	lgr         *golog.Core
	db          *mongo.Client
	pCollection *mongo.Collection
	read        *read
}

func (w *write) NewPost(title, body string, userId uint64, picturesNames []string) (string, error) {
	query := bson.M{
		"OwnerId":      userId,
		"Title":        title,
		"Body":         body,
		"PicturesName": picturesNames,
		"CreatedAt":    time.Now().String(),
	}

	result, err := w.pCollection.InsertOne(context.Background(), query)
	if err != nil {
		return "", err
	}

	objectId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("Cant convert object id")
	}

	return objectId.Hex(), nil
}

func (w *write) DeletePost(rawPostId string, userId uint64) (picturesName []string, err error) {
	postID, err := primitive.ObjectIDFromHex(rawPostId)
	if err != nil {
		return nil, err
	}

	query := bson.M{
		"_id":     postID,
		"OwnerId": userId,
	}

	result := w.pCollection.FindOneAndDelete(context.Background(), query)
	if err := result.Err(); err != nil {
		return nil, err
	}

	post := new(models.Post)

	return post.PicturesName, result.Decode(post)
}

func (w *write) DeleteUserPosts(userId uint64) ([]struct{ PicturesName []string }, error) {
	result := make([]struct {
		PicturesName []string
	}, 0)

	{
		query := bson.M{
			"OwnerId": userId,
		}

		cursor, err := w.pCollection.Find(context.Background(), query,
			options.Find().SetProjection(bson.M{"picturesName": 1}),
		)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(context.Background())

		err = cursor.All(context.Background(), &result)
		if err != nil {
			return nil, err
		}
	}

	go func() {
		query := bson.M{
			"OwnerId": userId,
		}

		_, err := w.pCollection.DeleteMany(context.Background(), query)
		if err != nil {
			panic(err)
		}
	}()

	return result, nil
}
