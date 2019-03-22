package model

import (
	"time"
)

type Post struct {
	Code        string    `bson:"code,omitempty"`
	Title       string    `bson:"title,omitempty"`
	Company     *Company  `bson:"company,omitempty"`
	Location    string    `bson:"location,omitempty"`
	Type        string    `bson:"type,omitempty"`
	Status      string    `bson:"status,omitempty"`
	Description string    `bson:"description,omitempty"`
	CreatedAt   time.Time `bson:"createdAt,omitempty"`
	UpdatedAt   time.Time `bson:"updatedAt,omitempty"`
	DeletedAt   time.Time `bson:"deletedAt,omitempty"`
}
