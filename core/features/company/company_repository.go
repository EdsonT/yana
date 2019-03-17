package company

import (
	"context"
	"log"
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
	coll = cli.Database("yanaDb").Collection("Company")
}

func AddPost(mpo *model.Company) (*mongo.InsertOneResult, error) {
	Init()
	return coll.InsertOne(context.TODO(), mpo)
}
func GetPost() []*model.Company {
	var results []*model.Company
	Init()

	cur, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem model.Company
		err = cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return results
}
