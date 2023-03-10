package controllers

import (
	"net/http"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/models"
	"github.com/labstack/echo/v4"
)

type categoryController struct {
	db databases.Database
}

func NewCategoryController(db databases.Database) *categoryController {
	return &categoryController{
		db: db,
	}
}

func (cc *categoryController) CreateCategory(c echo.Context) error {
	var category models.Category
	if c.Bind(&category) != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "invalid request"})
	} else if cc.db.CheckCategoryExists(category.Name) {
		return c.JSON(http.StatusConflict, echo.Map{"status": "category already exists"})
	} else {
		cc.db.CreateCategory(&category)
		return c.JSON(http.StatusOK, echo.Map{"status": "category created"})
	}
}

func (cc *categoryController) GetCategory(c echo.Context) error {
	category, _ := cc.db.GetCategory(c.Param("name"))
	if category.ID != 0 {
		return c.JSON(http.StatusOK, category)
	} else {
		return c.JSON(http.StatusNotFound, echo.Map{"status": "category not found"})
	}
}

func (cc *categoryController) GetCategories(c echo.Context) error {
	categories, _ := cc.db.GetCategories()
	return c.JSON(http.StatusOK, categories)
}
