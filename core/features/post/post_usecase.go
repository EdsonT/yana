package post

import (
	"time"
	"yana/model"

	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/mongo"
)

//CreateNewPost initializes primary parameters of a Post, and validate data
func CreateNewPost(params *model.Post) (*mongo.InsertOneResult, error) {
	np := new(model.Post)
	np.Code = xid.New().String()
	np.Title = params.Title
	np.Location = params.Location
	np.Company = params.Company
	np.Type = params.Type
	np.DateCreated = time.Now()
	np.LastUpdated = time.Now()
	result, err := AddPost(np)
	return result, err
}

func Search() {

}
