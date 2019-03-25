package post

import (
	"fmt"
	"log"
	"time"
	"yana/core/features/company"
	"yana/dao"
	"yana/errors"
	"yana/model"

	valid "github.com/asaskevich/govalidator"
	"github.com/rs/xid"
)

type PostService struct {
	errors errors.Error `json:"errors,omitempty"`
	Posts  interface{}  `json:"metadata,omitempty"`
}

//CreatePostImpl initializes primary parameters of a Post, and validate data
func (p *PostService) CreatePostImpl(params dao.Post) {
	var (
		np   model.Post
		errs errors.Error
	)

	// first validate if all the required params were provided
	valid.SetFieldsRequiredByDefault(true)
	_, err := valid.ValidateStruct(params)

	if err != nil {
		errs.SetValidationErrors(valid.ErrorsByField(err))
		p.errors = errs
		fmt.Println(errs)
	} else {
		// validate if the company provided is in the database, otherwise log and set the error
		if cp, err := company.Get(params.Company); cp != nil {
			fmt.Println(cp, err)
			np.Company = cp
			np.Code = xid.New().String()
			np.Status = model.PUBLISH
			np.Title = params.Title
			np.Type = params.Type
			np.Location = params.Location
			np.Description = params.Description
			np.CreatedAt = time.Now()
			// Insert document into the DB
			fmt.Println(np)
			inserted, err := Create(&np)
			if err != nil {
				errs.SetRepositoryError(valid.ErrorsByField(err))
				p.errors = errs
			}
			// Set Response
			p.Posts = &np
			log.Println("[LOG] Object inserted:", inserted.InsertedID)

		} else {
			p.errors.SetNotFoundOptionError(params.Company)
			log.Println(err)
		}

	}

}

//GetPostImpl retrieves the Posts with a simple creteria
func (p *PostService) GetPostListImpl(params dao.Post) {
	var pm model.Post
	pm.Code = params.Code
	pm.Title = params.Title
	pm.Status = params.Status
	pm.Location = params.Location

	//Get the company if provided
	if params.Company != "" {
		pm.Company, err = company.Get(params.Company)
	}
	pm.CreatedAt, _ = time.Parse(time.RFC3339, params.CreatedAt)

	postsFound := GetList(pm)
	//returns the all the posts in the response PostService
	p.Posts = postsFound
}

// func UpdatePost(code string, params model.Post) model.Post {
// 	var up model.Post
// 	Update(code, params)
// 	up = Get(code)
// 	return up
// }

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
