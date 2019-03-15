package post

import (
	"context"
	"fmt"
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
	coll = cli.Database("yanaDb").Collection("Posts")
}

func AddPost(mpo *model.Post) (*mongo.InsertOneResult, error) {
	Init()
	return coll.InsertOne(context.TODO(), mpo)
}
func GetPost() []*model.Post {
	var results []*model.Post
	Init()
	// findOpts := options.Find()
	// findOpts.SetLimit(3)
	cur, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem model.Post
		err = cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)

	}
	fmt.Println(results)

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return results
}

func FindPost(id string){
	Init()
	res:=coll.FindOne(context.TODO(),id)
	fmt.Println(res)
}
