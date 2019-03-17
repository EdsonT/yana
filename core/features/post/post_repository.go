package post

import (
	"context"
	"fmt"
	"log"
	"yana/config"
	"yana/model"

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
func GetPost(params model.Post) []*model.Post {
	var (
		results []*model.Post
		cur     *mongo.Cursor
		err     error
	)
	Init()

	cur, err = coll.Find(context.TODO(), params)

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

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return results
}

func FindPost(id string) {
	Init()
	res := coll.FindOne(context.TODO(), id)
	fmt.Println(res)
}
func CountRecords() int64 {
	Init()
	// var opts *options.EstimatedDocumentCountOptions
	// opts.SetMaxTime(3)
	res, _ := coll.EstimatedDocumentCount(context.Background())
	fmt.Println(res)
	return res
}
