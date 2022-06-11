package controllers

import (
	"net/http"

	"github.com/arshamalh/blogo/database"
	"github.com/arshamalh/blogo/models"
	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category models.Category
	if c.BindJSON(&category) == nil {
		if !database.CheckCategoryExists(category.Name) {
			database.CreateCategory(&category)
			c.JSON(http.StatusOK, gin.H{"status": "category created"})
		} else {
			c.JSON(http.StatusConflict, gin.H{"status": "category already exists"})
		}
	}
}

func GetCategory(c *gin.Context) {
	category, _ := database.GetCategory(c.Param("name"))
	if category.ID != 0 {
		c.JSON(http.StatusOK, category)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": "category not found"})
	}
}

func GetCategories(c *gin.Context) {
	categories, _ := database.GetCategories()
	c.JSON(http.StatusOK, categories)
}
