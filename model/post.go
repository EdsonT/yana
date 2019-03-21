package model

import (
	"time"
)

type Post struct {
	Code        string    `form:"code" bson:"code,omitempty"`
	Title       string    `form:"title" bson:"title,omitempty"`
	Company     string    `form:"company" bson:"company,omitempty"`
	Location    string    `form:"location" bson:"location,omitempty"`
	Type        string    `form:"type" bson:"type,omitempty"`
	Status      string    `form:"status" bson:"status,omitempty"`
	Description string    `form:"description" bson:"description,omitempty"`
	CreatedAt   time.Time `form:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt   time.Time `form:"updatedAt" bson:"updatedAt,omitempty"`
	DeletedAt   time.Time `form:"deletedAt" bson:"deletedAt,omitempty"`
}
