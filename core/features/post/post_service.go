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
func CreatePostImpl(params dao.Post) (model.Post, error) {
	np := new(model.Post)
	np.Code = xid.New().String()
	np.Title = params.Title
	np.Location = params.Location
	cp, err := company.Find(params.Company)
	if cp == (model.Company{}) {

		return model.Post{}, err

	}
	np.Company = cp
	np.Type = params.Type
	np.Status = "Active"
	np.CreatedAt = time.Now()
	np.UpdatedAt = time.Now()
	result, err := Add(np)
	log.Println("Object Inserted:", result)
	return Find(np.Code), err
}

func GetPostImpl(params dao.Post) []*model.Post {
	// var listPosts []*model.Post
	postsFound := Get(params)
	// for _, post := range postsFound {
	// 	np := new(dao.Post)
	// 	np.Code = post.Code
	// 	np.Title = post.Title
	// 	np.Status = post.Status
	// 	np.Location = post.Location
	// 	np.Company, err = company.Find(post.Company.Code)
	// 	listPosts = append(listPosts, np)
	// }

	return postsFound
}
func UpdatePost(code string, params model.Post) model.Post {
	var up model.Post
	Update(code, params)
	up = Find(code)
	return up
}

// // func SearchPosts(keyword string) []model.Post {
// // 	var result []model.Post
// // 	return result
// // }
// // func DeletePost(code string) (string, error) {
// // 	result, err := DeleteLogical(code)
// // 	if err == nil {
// // 		log.Println("Object Deleted:", result)
// // 		return "success", nil
// // 	} else {
// // 		log.Println(err)
// // 		return "", err
// // 	}

// }
