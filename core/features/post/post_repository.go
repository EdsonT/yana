package post

import (
	"context"
	"fmt"
	"log"
	"yana/config"
	"yana/dao"
	"yana/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func Add(mpo *model.Post) (*mongo.InsertOneResult, error) {
	Init()
	return coll.InsertOne(context.TODO(), mpo)
}
func Get(params dao.Post) []*model.Post {
	var (
		results []*model.Post
		cur     *mongo.Cursor
		err     error
	)
	Init()
	json, err := bson.Marshal(params)
	fmt.Println(json)
	cur, err = coll.Find(context.TODO(), json)

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
func Update(code string, params model.Post) (*mongo.UpdateResult, error) {
	Init()
	filter := bson.D{{"code", code}}
	update := bson.D{
		{"$set", params},
	}
	res, err := coll.UpdateOne(context.TODO(), filter, update)
	fmt.Println(res)
	return res, err
}

func Find(code string) model.Post {
	Init()
	var res model.Post
	filter := bson.D{{"code", code}}
	coll.FindOne(context.TODO(), filter).Decode(&res)

	return res
}
func CountRecords() int64 {
	Init()
	// var opts *options.EstimatedDocumentCountOptions
	// opts.SetMaxTime(3)
	res, _ := coll.EstimatedDocumentCount(context.Background())
	fmt.Println(res)
	return res
}
func Search(kw string) []*model.Post {
	Init()
	var (
		res []*model.Post
		cur *mongo.Cursor
		err error
	)
	skey := bson.M{"title": primitive.Regex{Pattern: ".*" + kw + ".*", Options: "i"}}
	cur, err = coll.Find(context.Background(), skey)

	for cur.Next(context.TODO()) {
		var elem model.Post
		err = cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, &elem)
		fmt.Println(res)
	}
	return res
}
func DeletePostPhysic(cod string) *mongo.SingleResult {
	Init()
	res := coll.FindOneAndDelete(context.TODO(), bson.D{{"code", cod}})
	return res
}
func DeleteLogical(cod string) (*mongo.UpdateResult, error) {
	Init()
	params := bson.D{{"status", "Inactive"}}
	filter := bson.D{{"code", cod}}
	update := bson.D{
		{"$set", params},
	}
	res, err := coll.UpdateOne(context.TODO(), filter, update)
	return res, err
}
