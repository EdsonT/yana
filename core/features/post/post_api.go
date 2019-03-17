package post

import (
	// "fmt"

	"fmt"
	"log"
	"yana/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func SetBaseRoutePost(r *gin.Engine) *gin.Engine {
	rg := r.Group("posts")
	AddRoutes(rg)
	return r
}

func AddRoutes(rg *gin.RouterGroup) {

	rg.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	rg.GET("/", func(c *gin.Context) {
		var np model.Post

		if c.ShouldBind(&np) == nil {
			log.Println(np)
		}
		CountRecords()
		c.JSON(200, GetPosts(np))

	})
	rg.POST("/create", func(c *gin.Context) {
		var np *model.Post
		c.BindJSON(&np)
		result, _ := CreateNewPost(np)
		c.JSON(200, result)

	})
	rg.PUT("/:code", func(c *gin.Context) {
		var np model.Post

		c.BindJSON(&np)
		fmt.Println(c.Param("code"))
		res := UpdatePost(c.Param("code"), np)
		c.JSON(200, res)

	})
	// rg.DELETE("/:code", func(c *gin.Context) {
	// 	code := c.Param("code")
	// 	DeletePostPhysics(code)
	// 	c.JSON(200, "success")
	// })
	rg.DELETE("/:code", func(c *gin.Context) {
		res, err := DeletePost(c.Param("code"))
		if res != "" {
			c.JSON(200, bson.M{"status": res})
		} else {
			c.JSON(500, err)
		}
	})

}
