package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//Return the post

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	//Get the post
	var posts []models.Post
	initializers.DB.Find(&posts)
	//Respond with them

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	//Get the post
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}
