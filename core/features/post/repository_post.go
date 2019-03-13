package post

import (
	"context"
	"time"
	"yana/config"
	"yana/model"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	cli  *mongo.Client
	err  error
	coll *mongo.Collection
)

func Init() {
	cli = config.NewMongoClient()
	coll = cli.Database("yanaDb").Collection("Posts")
}

func CreatePost(mpo model.Post) {
	Init()
	coll.InsertOne(context.TODO(), bson.M{
		"code":        mpo.Code,
		"title":       mpo.Title,
		"company":     mpo.Company,
		"location":    mpo.Location,
		"type":        mpo.Type,
		"description": mpo.Description,
		"dateCreated": time.Now(),
		"lastUpdated": time.Now(),
	})
}
