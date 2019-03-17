package post

import (
	"context"
	"testing"
	"time"
	"yana/model"

	"github.com/rs/xid"

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
		t.Fatal("Test Failed", params.Code, retr.Code)
		t.Fail()
	}

}

func TestSearch(t *testing.T) {
	reg := new(model.Post)
	reg.Company = "testCompany"
	reg.Title = "testTitle"
	reg.Type = "testType"
	reg.Location = "testLocation"
	CreateNewPost(reg)
	var params *model.Post
	params = new(model.Post)
	params.Title = "testTitle"
	result := Search(params)

	if params == nil {
		total := CountRecords()
		if len(result) != total {
			t.Fatal("Test Failed")
		}
	} else {
		if len(result) == 0 {
			t.Fatal("Test Failed")
		}
	}

	params.Code = xid.New().String()

	result := Search(params)

	if result > 0 {
		t.Fatal("Test Failed")
	}

}
