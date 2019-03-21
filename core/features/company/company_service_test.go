package company

import (
	"context"
	"testing"
	"time"
	"yana/model"

	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateNewPost(t *testing.T) {
	params := new(model.Company)
	params.Code = "TestingCodeQWERTY"
	result, err := CreateNewPost(params)
	if err != nil {
		t.Fatal(err)
	}
	filter := bson.D{{"_id", result.InsertedID}}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var retr model.Company
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
	reg := new(model.Company)
	reg.Name = "testName"
	reg.City = "testCity"
	reg.Country = "testCountry"
	reg.Address = "testAddress"
	CreateNewPost(reg)
	var params *model.Company
	params = new(model.Company)
	params.Name = "testName"
	result := Search()

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

	result := Search()

	if result > 0 {
		t.Fatal("Test Failed")
	}

}
