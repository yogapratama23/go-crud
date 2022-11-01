package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yogapratama23/go-crud/initializers"
	"github.com/yogapratama23/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		Title string
		Description string
	}
	c.Bind(&body)

	post := models.Post{ Title: body.Title, Description: body.Description }
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400);
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post created!",
		"data": post,
	})
}

func PostsList(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		c.Status(404)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Post list!",
		"data": posts,
	})
}

func PostDetails(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Post detail!",
		"data": post,
	})
}

func PostUpdate(c *gin.Context) {
	var body struct {
		Title string
		Description string
	}
	c.Bind(&body)
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)
	initializers.DB.Model(&post).Updates(models.Post {
		Title: body.Title,
		Description: body.Description,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated!",
		"data": post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted!",
		"data": "",
	})
}
