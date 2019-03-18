package company

import (
	"fmt"
	"log"
	"yana/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
		var op *model.Company
		c.BindJSON(&np)
		// fmt.Println(c.Param("code"))
		result, _ := CreateNewCompany(np)
		fmt.Println(result)
		op = FindC(np.Name)
		c.JSON(200, op)
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
