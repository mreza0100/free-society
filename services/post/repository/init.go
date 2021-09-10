package repository

import (
	"context"
	"fmt"
	"freeSociety/services/post/configs"
	"freeSociety/services/post/instances"
	"time"

	"github.com/mreza0100/golog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewRepo(lgr *golog.Core) *instances.Repository {
	var (
		db             *mongo.Client
		postCollection *mongo.Collection
		readQ          *read
		writeQ         *write
	)

	{
		db = getConnection(lgr)
		postCollection = db.Database("posts").Collection("posts")
		lgr = lgr.With("In Repository ->")
	}
	{
		readQ = &read{
			lgr:         lgr.With("In Read ->"),
			db:          db,
			pCollection: postCollection,
		}
		writeQ = &write{
			lgr:         lgr.With("In Read ->"),
			db:          db,
			pCollection: postCollection,
		}
		readQ.write = writeQ
		writeQ.read = readQ
	}

	return &instances.Repository{
		Read:  readQ,
		Write: writeQ,
	}
}

func getConnection(lgr *golog.Core) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	mongoPort := configs.Configs.Mongo_port
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(fmt.Sprintf("mongodb://localhost:%v", mongoPort)),
		options.Client().SetMaxPoolSize(10),
	)
	if err != nil {
		lgr.Fatal("mongo connection failed ", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		lgr.Fatal("mongo ping failed ", err)
	}

	return client
}
