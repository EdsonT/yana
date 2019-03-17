package company

import (
	"context"
	"fmt"
	"log"
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
	coll = cli.Database("yanaDb").Collection("Company")
}

func AddPost(mpo *model.Company) (*mongo.InsertOneResult, error) {
	Init()
	return coll.InsertOne(context.TODO(), mpo)
}
func GetPost(params model.Company) []*model.Company {
	var (
		results []*model.Company
		cur     *mongo.Cursor
		err     error
	)
	Init()

	cur, err = coll.Find(context.TODO(), params)

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

func DeletCompanyPhysics(cod string) *mongo.SingleResult {
	Init()
	res := coll.FindOneAndDelete(context.TODO(), bson.D{{"code", cod}})
	return res
}
func DeletCompanyLogical(cod string) (*mongo.UpdateResult, error) {
	Init()
	params := bson.D{{"status", "Inactive"}}
	filter := bson.D{{"code", cod}}
	update := bson.D{
		{"$set", params},
	}
	res, err := coll.UpdateOne(context.TODO(), filter, update)
	fmt.Println(res, err)
	return res, err
}
func UpdatePost(code string, params model.Company) (*mongo.UpdateResult, error) {
	Init()
	filter := bson.D{{"code", code}}
	params.LastUpdated = time.Now()
	update := bson.D{
		{"$set", params},
	}
	res, err := coll.UpdateOne(context.TODO(), filter, update)
	// updatwe := bson.D{
	// 	{"$set", bson.D{"lastupdated", time.Now()}},
	// }
	// res, err := coll.UpdateOne(context.TODO(), filter, update)
	fmt.Println(res)
	return res, err

}
