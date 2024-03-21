package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	m "GIN/model"
)

func CreateItem(c *gin.Context) {
	var newItem m.Item
	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save newItem to the database (e.g., using ORM or database driver)
	// Example:
	db.Create(&newItem)

	c.JSON(http.StatusCreated, newItem)
}

func GetItems(c *gin.Context) {
	// Retrieve items from the database
	var items []m.Item
	db.Find(&items)

	c.JSON(http.StatusOK, items)
}

func GetItemByID(c *gin.Context) {
	id := c.Param("id")
	// Retrieve item from the database based on id
	// Example:
	var item m.Item
	db.First(&item, "id = ?", id)

	// Check if item exists
	if item.ID == "0" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var updatedItem m.Item
	if err := c.BindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update item in the database based on id
	// Example:
	db.Model(&m.Item{}).Where("id = ?", id).Updates(updatedItem)

	c.JSON(http.StatusOK, updatedItem)
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	// Delete item from the database based on id
	// Example:
	db.Where("id = ?", id).Delete(&m.Item{})

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
}
