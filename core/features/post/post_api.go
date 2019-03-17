package post

import (
	// "fmt"

	"fmt"
	"log"
	"yana/model"

	"github.com/gin-gonic/gin"
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

	rg.POST("/create", func(c *gin.Context) {
		var np *model.Post
		c.BindJSON(&np)
		CreateNewPost(np)
		c.JSON(200, "success")

	})
	rg.GET("/", func(c *gin.Context) {
		var np model.Post

		if c.ShouldBind(&np) == nil {
			log.Println(np)
		}
		// fmt.Println(np)
		CountRecords()
		c.JSON(200, GetPost(np))

	})

	rg.PUT("/:code", func(c *gin.Context) {
		var np model.Post

		c.BindJSON(&np)
		fmt.Println(c.Param("code"))
		res,_ := UpdatePost(c.Param("code"), np)
		c.JSON(200, res)

	})
	// rg.DELETE("/:code", func(c *gin.Context) {
	// 	code := c.Param("code")
	// 	DeletePostPhysics(code)
	// 	c.JSON(200, "success")
	// })
	rg.DELETE("/:code", func(c *gin.Context) {
		code := c.Param("code")
		DeletePostLogical(code)
		c.JSON(200, "success")
	})

}
