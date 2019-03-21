package post

import (
	"log"
	"time"
	"yana/core/features/company"
	"yana/dao"
	"yana/model"

	"github.com/rs/xid"
)

//CreatePostImpl  initializes primary parameters of a Post, and validate data
func CreatePostImpl(params *model.Post) (model.Post, error) {
	np := new(model.Post)
	np.Code = xid.New().String()
	np.Title = params.Title
	np.Location = params.Location
	np.Company = params.Company
	np.Type = params.Type
	np.Status = "Active"
	np.CreatedAt = time.Now()
	np.UpdatedAt = time.Now()
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
func GetPostImpl(params model.Post) []*dao.Post {
	var listPosts []*dao.Post
	postsFound := Get(params)
	for _, post := range postsFound {
		np := new(dao.Post)
		np.Code = post.Code
		np.Title = post.Title
		np.Status = post.Status
		np.Location = post.Location
		np.Company = company.Find(post.Company)
		listPosts = append(listPosts, np)
	}

	return listPosts
}

func SearchPosts(keyword string) []model.Post {
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
