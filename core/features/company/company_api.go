package company

import (
	"fmt"
	"log"
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
		var np model.Company

		if c.ShouldBind(&np) == nil {
			log.Println(np)
		}
		fmt.Println(np)
		CountRecords()
		c.JSON(200, GetPost(np))

	})
	rg.DELETE("/:code", func(c *gin.Context) {
		code := c.Param("code")
		DeletCompanyPhysics(code)
		c.JSON(200, "success")
	})
	rg.PUT("/:code", func(c *gin.Context) {
		code := c.Param("code")
		DeletCompanyLogical(code)
		c.JSON(200, "success")
	})
}
