package router

import (
	"yana/core/features/company"
	"yana/core/features/post"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//NewRouter will create create a gin router and return a gin.engine
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           86400,
	}))
	// config.AllowOrigins = []string{"http://localhost:80"}
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}

	// r.Use(cors.New(config))

	InitRoutes(r)
	return r
}

//InitRoutes registers all the API endpoints
func InitRoutes(app *gin.Engine) {
	app.GET("/", post.GetPostListController)
	app.OPTIONS("", func(g *gin.Context) {
		g.JSON(200, gin.H{"foo": "bar"})
	})
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
