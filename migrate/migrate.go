package main

import (
	"github.com/yogapratama23/go-crud/initializers"
	"github.com/yogapratama23/go-crud/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
