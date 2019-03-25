package post

import (
	"context"
	"fmt"
	"log"
	"yana/config"
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

//Init intializes the conection with the database
func Init() {
	cli = config.NewMongoClient()
	coll = cli.Database("yanaDB").Collection("Posts")
}

//Create inserts a new Post document into the database
func Create(mpo *model.Post) (*mongo.InsertOneResult, error) {
	Init()
	return coll.InsertOne(context.TODO(), mpo)
}

//GetList retrieves the posts given a criteria based on the Post model
func GetList(params model.Post) []*model.Post {
	var (
		results []*model.Post
		cur     *mongo.Cursor
	)
	Init()
	cur, err = coll.Find(context.TODO(), params)
	//If error encountered when retrieven documents
	if err != nil {
		log.Println(err)
	}
	for cur.Next(context.TODO()) {
		var elem model.Post
		err = cur.Decode(&elem)
		if err != nil {
			log.Println(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
	}
	return results
}

//Update function update a Post with a given valid Post Model
func Update(code string, params model.Post) (*mongo.UpdateResult, error) {
	Init()
	filter := bson.D{{"code", code}}
	update := bson.D{
		{"$set", params},
	}
	res, err := coll.UpdateOne(context.TODO(), filter, update)
	return res, err
}

// Get one Post by its code
func Get(code string) model.Post {
	Init()
	var res model.Post
	filter := bson.D{{"code", code}}
	coll.FindOne(context.TODO(), filter).Decode(&res)

	return res
}

// CountRecords counts the number of documents
func CountRecords() int64 {
	Init()
	res, _ := coll.EstimatedDocumentCount(context.Background())
	return res
}

// Search returns a list of records based on a Search model
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

//DeletePhysical will remove the document from the database
func DeletePhysical(cod string) *mongo.SingleResult {
	Init()
	res := coll.FindOneAndDelete(context.TODO(), bson.D{{"code", cod}})
	return res
}

//DeleteLogical will deactivate the document
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
