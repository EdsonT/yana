package main

import (
	// "context"

	"fmt"

	"yana/core/features/post"

	"github.com/gin-gonic/gin"
	// "yana/config"
	// "go.mongodb.org/mongo-driver/bson"
)

func main() {
	fmt.Println("Working!")
	r := gin.Default()
	post.SetBaseRoutePost(r)
	r.Run()
}
