package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.POST("/items", CreateItem)
	r.GET("/items", GetItems)
	r.GET("/items/:id", GetItemByID)
	r.PUT("/items/:id", UpdateItem)
	r.DELETE("/items/:id", DeleteItem)

	return r
}
