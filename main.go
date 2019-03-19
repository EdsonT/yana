package main

import (
	"yana/core/features/company"
	"yana/core/features/post"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	post.SetBaseRoutePost(r)
	company.SetBaseRoutePost(r)
	r.Run()
}
