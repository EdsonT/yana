package dao

type Post struct {
	Code        string `form:"code,omitempty" bson:"code,omitempty"`
	Title       string `form:"title,omitempty" bson:"title,omitempty"`
	Company     string `form:"company,omitempty" bson:"company,omitempty"`
	Location    string `form:"location,omitempty" bson:"location,omitempty"`
	Type        string `form:"type,omitempty" bson:"type,omitempty"`
	Status      string `form:"status,omitempty" bson:"status,omitempty"`
	Description string `form:"description,omitempty" bson:"description,omitempty"`
	CreatedAt   string `form:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   string `form:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	DeletedAt   string `form:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
}
