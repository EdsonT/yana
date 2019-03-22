package company

import (
	"fmt"
	"log"
	"yana/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateCompanyController(c *gin.Context) {
	var np *model.Company
	var result *model.Company
	c.BindJSON(&np)
	result = CreateNewCompany(np)
	c.JSON(200, result)
}
func GetCompanyController(c *gin.Context) {
	var np model.Company

	if c.ShouldBind(&np) == nil {
		log.Println(np)
	}
	c.JSON(200, GetCompany(np))
}

func AddRoutes(rg *gin.RouterGroup) {

	rg.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	rg.GET("/", func(c *gin.Context) {
		var np model.Company

		if c.ShouldBind(&np) == nil {
			log.Println(np)
		}
		CountRecords()
		c.JSON(200, GetCompany(np))

	})

	rg.POST("/create", func(c *gin.Context) {
		var np *model.Company
		var result *model.Company
		c.BindJSON(&np)
		result = CreateNewCompany(np)
		c.JSON(200, result)
	})

	rg.PUT("/:code", func(c *gin.Context) {
		var np model.Company
		c.BindJSON(&np)
		fmt.Println(c.Param("code"))
		res := UpdateCompany(c.Param("code"), np)

		c.JSON(200, res)
	})

	// rg.DELETE("/:code", func(c *gin.Context) {
	// 	code := c.Param("code")
	// 	DeletCompanyPhysics(code)
	// 	c.JSON(200, "success")
	// })
	rg.DELETE("/:code", func(c *gin.Context) {
		res, err := Delete(c.Param("code"))
		if res != "" {
			c.JSON(200, bson.M{"status": res})
		} else {
			c.JSON(500, err)
		}
	})

}
