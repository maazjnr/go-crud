package main

import (
	"fmt"
	"go-crud/controllers"
	"go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostIndex)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.GET("/posts/:id", controllers.PostShow)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.Run()

	fmt.Println("Hello world")
}
