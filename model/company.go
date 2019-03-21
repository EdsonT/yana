package model

import "time"

type Company struct {
	Code        string    `form:"code,omitempty" json:"code,omitempty" bson:"code,omitempty"`
	Name        string    `form:"name,omitempty" json:"name,omitempty" bson:"name,omitempty"`
	City        string    `form:"city,omitempty" json:"city,omitempty" bson:"city,omitempty"`
	Country     string    `form:"country,omitempty" json:"country,omitempty" bson:"country,omitempty"`
	Address     string    `form:"address,omitempty" json:"address,omitempty" bson:"address,omitempty"`
	Status      string    `form:"status,omitempty" json:"status,omitempty" bson:"status,omitempty"`
	Description string    `form:"description,omitempty" json:"description,omitempty" bson:"description,omitempty"`
	CreatedAt   time.Time `form:"createdAt,omitempty" json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   time.Time `form:"updatedAt,omitempty" json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	DeletedAt   time.Time `form:"deletedAt,omitempty" json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
}
