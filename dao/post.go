package dao

import (
	"yana/model"
)



type Post struct {
	Code        string `bson:"code"`
	Title       string `bson:"title"`
	Company     model.Company `bson:"company"`
	Location    string `bson:"location"`
	Type        string `bson:"type"`
	Status      string `bson:"status"`
	Description string `bson:"description"`
	CreatedAt   string `bson:"createdAt"`
	UpdatedAt   string `bson:"updatedAt"`
	DeletedAt   string `bson:"deletedAt"`
}
