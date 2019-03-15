package post

import (
	"context"
	"testing"
	"time"
	"yana/model"

	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateNewPost(t *testing.T) {
	params := new(model.Post)
	params.Code = "TestingCodeQWERTY"
	result, err := CreateNewPost(params)
	if err != nil {
		t.Fatal(err)
	}
	filter := bson.D{{"_id", result.InsertedID}}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var retr model.Post
	err = coll.FindOne(ctx, filter).Decode(&retr)
	if err != nil {
		t.Fatal(err)
	}
	if params.Code != retr.Code {
		t.Fatal("Test Failed", params.Code,retr.Code)
		t.Fail()
	}

}
