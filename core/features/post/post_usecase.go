package post

import (
	"log"
	"time"
	"yana/model"

	"github.com/rs/xid"
)

//CreateNewPost initializes primary parameters of a Post, and validate data
func CreateNewPost(params *model.Post) (model.Post, error) {
	np := new(model.Post)
	np.Code = xid.New().String()
	np.Title = params.Title
	np.Location = params.Location
	np.Company = params.Company
	np.Type = params.Type
	np.Status = "Active"
	np.DateCreated = time.Now()
	np.LastUpdated = time.Now()
	result, err := Add(np)
	log.Println("Object Inserted:", result)
	return Find(np.Code), err
}
func UpdatePost(code string, params model.Post) model.Post {
	var up model.Post
	Update(code, params)
	up = Find(code)
	return up
}
func GetPosts(params model.Post) []*model.Post {
	return Get(params)
}

func Search(params *model.Post) []model.Post {
	var result []model.Post
	return result
}
func DeletePost(code string) (string, error) {
	result, err := DeleteLogical(code)
	if err == nil {
		log.Println("Object Deleted:", result)
		return "success", nil
	} else {
		log.Println(err)
		return "", err
	}

}
