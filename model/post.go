package model

import (
	"time"
)

type Post struct {
	Code        string    `form:"code,omitempty" json:"code,omitempty" bson:"code,omitempty"`
	Title       string    `form:"title,omitempty" json:"title,omitempty" bson:"title,omitempty"`
	Company     string    `form:"company,omitempty" json:"company,omitempty" bson:"company,omitempty"`
	Location    string    `form:"location,omitempty" json:"location,omitempty" bson:"location,omitempty"`
	Type        string    `form:"type,omitempty" json:"type,omitempty" bson:"type,omitempty"`
	Status      string    `form:"status,omitempty" json:"status,omitempty" bson:"status,omitempty"`
	Description string    `form:"description,omitempty" json:"description,omitempty" bson:"description,omitempty"`
	DateCreated time.Time `form:"dateCreated,omitempty" json:"dateCreated,omitempty" bson:"dateCreated,omitempty"`
	LastUpdated time.Time `form:"lastUpdated,omitempty" json:"lastUpdated,omitempty" bson:"lastUpdated,omitempty"`
}
