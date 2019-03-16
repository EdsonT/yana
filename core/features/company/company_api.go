package company

import (
	// "fmt"

	"yana/model"

	"github.com/gin-gonic/gin"
)

func SetBaseRoutePost(r *gin.Engine) *gin.Engine {
	rg := r.Group("company")
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
		var np *model.Company
		c.BindJSON(&np)
		CreateNewPost(np)
		c.JSON(200, "success")

	})
	rg.GET("/", func(c *gin.Context) {
		c.JSON(200, GetPost())
	})

}
