package router

import (
	"yana/core/features/company"
	"yana/core/features/post"

	"github.com/gin-gonic/gin"
)

//NewRouter will create create a gin router and return a gin.engine
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	InitRoutes(r)
	return r
}

//InitRoutes registers all the API endpoints
func InitRoutes(app *gin.Engine) {

	posts := app.Group("posts")
	{
		posts.GET("/", post.GetPostListController)
		posts.GET("/search")
		posts.POST("/create", post.CreatePostController)
		posts.PUT("/:code/update")
		posts.DELETE("/:code/delete")
	}
	companies := app.Group("companies")
	{
		companies.GET("/", company.GetCompanyController)
		companies.POST("/create", company.CreateCompanyController)
		companies.PUT("/:code/update")
		companies.DELETE("/:code/delete")
	}

}
