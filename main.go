package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yogapratama23/go-crud/controllers"
	"github.com/yogapratama23/go-crud/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default();
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		"message": "default endpoint",
		});
	});
	r.GET("/posts", controllers.PostsList)
	r.POST("/posts/create", controllers.PostsCreate)
	r.POST("/posts/:id", controllers.PostDetails)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.Run();
}
