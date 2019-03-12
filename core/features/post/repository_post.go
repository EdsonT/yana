package post

import (
	"context"
	"time"
	"yana/config"
	"yana/model"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

var coll *mongo.Collection

func initCollection() {
	coll = config.GetClient().Database("yanaDB").Collection("posts")

}

func CreatePost(mpo model.Post) {
	initCollection()
	coll.InsertOne(context.TODO(), bson.M{
		"code":        mpo.Code,
		"title":       mpo.Title,
		"company":     mpo.Company,
		"location":    mpo.Location,
		"description": mpo.Description,
		"dateCreated": time.Now(),
		"lastUpdated": time.Now(),
	})
}
