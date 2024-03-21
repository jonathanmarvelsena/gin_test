package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Endpoints
	router.POST("/items", CreateItem)
	router.GET("/items", GetItems)
	router.GET("/items/:id", GetItemByID)
	router.PUT("/items/:id", UpdateItem)
	router.DELETE("/items/:id", DeleteItem)

	router.Run(":8080")
}
