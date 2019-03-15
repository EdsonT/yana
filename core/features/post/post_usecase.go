package post

import (
	"time"
	"yana/model"

	"go.mongodb.org/mongo-driver/mongo"
)
//CreateNewPost initializes primary parameters of a Post, and validate data
func CreateNewPost(params *model.Post) (*mongo.InsertOneResult, error) {
	np := new(model.Post)
	np.Code=params.Code
	np.Title=params.Title
	np.Location=params.Location
	np.Company=params.Company
	np.Type=params.Type
	np.DateCreated = time.Now()
	np.LastUpdated = time.Now()
	result, err := AddPost(np)
	return result, err
}
