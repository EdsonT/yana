package dao

type Post struct {
	Code        string `form:"code,omitempty" bson:"code,omitempty" valid:"-"`
	Title       string `form:"title,omitempty" bson:"title,omitempty" valid:"-"`
	Company     string `form:"company," bson:"company,omitempty" valid:"-"`
	Location    string `form:"location,omitempty" bson:"location,omitempty" valid:"-"`
	Type        string `form:"type,omitempty" bson:"type,omitempty" valid:"required~Value required"`
	Status      string `form:"status,omitempty" bson:"status,omitempty" valid:"required"`
	Description string `form:"description,omitempty" bson:"description,omitempty" valid:"optional"`
	CreatedAt   string `form:"createdAt,omitempty" bson:"createdAt,omitempty" valid:"optional"`
	UpdatedAt   string `form:"updatedAt,omitempty" bson:"updatedAt,omitempty" valid:"optional"`
	DeletedAt   string `form:"deletedAt,omitempty" bson:"deletedAt,omitempty" valid:"optional"`
}

// func (dao *Post) GetPostModel() (model.Post, error) {
// 	np := model.Post{}
// 	nc, err := company.Find(dao.Company)
// 	if (err != nil) || (nc != nil) {
// 		np.Code = dao.Code
// 		np.Title = dao.Title
// 		np.Company = nc
// 		np.Location = dao.Location
// 		np.Type = dao.Type
// 		np.Status = dao.Status
// 		np.Description = dao.Description
// 		np.CreatedAt, _ = time.Parse(time.RFC3339, dao.CreatedAt)
// 		np.UpdatedAt, _ = time.Parse(time.RFC3339, dao.UpdatedAt)
// 		np.DeletedAt, _ = time.Parse(time.RFC3339, dao.DeletedAt)
// 		return np, err
// 	} else {
// 		return np, errors.New("Invalid Company")
// 	}
// 	return np, err

// }
