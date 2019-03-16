package company

import (
	"time"
	"yana/model"

	"go.mongodb.org/mongo-driver/mongo"
)

//CreateNewPost initializes primary parameters of a Post, and validate data
func CreateNewPost(params *model.Company) (*mongo.InsertOneResult, error) {
	np := new(model.Company)
	np.Code = params.Code
	np.Name = params.Name
	np.City = params.City
	np.Country = params.Country
	np.Description = params.Description
	np.DateCreated = time.Now()
	np.LastUpdated = time.Now()
	result, err := AddPost(np)
	return result, err
}
