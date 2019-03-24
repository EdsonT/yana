package post

import (
	valid "github.com/asaskevich/govalidator"
	// "fmt"

	"fmt"
	"log"
	"yana/dao"
	"yana/errors"
	"yana/model"

	"github.com/gin-gonic/gin"
)

func GetPostController(c *gin.Context) {
	var params dao.Post

	if c.ShouldBind(&params) == nil {
		log.Println(params)
	}
	c.JSON(200, GetPostImpl(params))
}
func CreatePostController(c *gin.Context) {
	valid.SetFieldsRequiredByDefault(true)
	var params dao.Post
	c.ShouldBind(&params)
	//get validation
	validation, err := valid.ValidateStruct(params)
	if err != nil {
		var e errors.Error
		e.SetValidationErrors(valid.ErrorsByField(err))
		c.JSON(400, e)
		return
	}
	fmt.Println(validation)
	res, _ := CreatePostImpl(params)
	c.JSON(200, res)

}

func AddRoutes(rg *gin.RouterGroup) {

	rg.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// rg.GET("/", func(c *gin.Context) {
	// 	var np model.Post

	// 	if c.ShouldBind(&np) == nil {
	// 		log.Println(np)
	// 	}
	// 	CountRecords()
	// 	c.JSON(200, GetPostsImpl(np))

	// })
	rg.GET("/search", func(c *gin.Context) {
		type SearchDAO struct {
			Keyword string `form:"keyword" json:"keyword" bson:"keyword"`
		}
		var (
			sd      SearchDAO
			keyword string
		)
		if c.ShouldBind(&sd) == nil {
			keyword = sd.Keyword
		}

		c.JSON(200, Search(keyword))
	})
	// rg.POST("/create", func(c *gin.Context) {
	// 	var np *model.Post
	// 	c.BindJSON(&np)
	// 	result, _ := CreateNewPost(np)
	// 	c.JSON(200, result)

	// })
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
	// rg.DELETE("/:code", func(c *gin.Context) {
	// 	res, err := DeletePost(c.Param("code"))
	// 	if res != "" {
	// 		c.JSON(200, bson.M{"status": res})
	// 	} else {
	// 		c.JSON(500, err)
	// 	}
	// })

}
