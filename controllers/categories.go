package controllers

import (
	"net/http"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/models"
	"github.com/gin-gonic/gin"
)

type categoryController struct {
	db databases.Database
}

func NewCategoryController(db databases.Database) *categoryController {
	return &categoryController{
		db: db,
	}
}

func (cc *categoryController) CreateCategory(c *gin.Context) {
	var category models.Category
	if c.BindJSON(&category) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "invalid request"})
	} else if cc.db.CheckCategoryExists(category.Name) {
		c.JSON(http.StatusConflict, gin.H{"status": "category already exists"})
	} else {
		cc.db.CreateCategory(&category)
		c.JSON(http.StatusOK, gin.H{"status": "category created"})
	}
}

func (cc *categoryController) GetCategory(c *gin.Context) {
	category, _ := cc.db.GetCategory(c.Param("name"))
	if category.ID != 0 {
		c.JSON(http.StatusOK, category)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": "category not found"})
	}
}

func (cc *categoryController) GetCategories(c *gin.Context) {
	categories, _ := cc.db.GetCategories()
	c.JSON(http.StatusOK, categories)
}
