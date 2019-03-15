package post

import (
	"yana/model"
)

func CreateNewPostTest() {
	params := new(model.Post)
	result, err := CreateNewPost(params)
	var id string
	id = result.InsertedID.str
	//look for document
	FindPost(result.InsertedID)
}
