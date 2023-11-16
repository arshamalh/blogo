package controllers

import (
	"net/http"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/log"
	"github.com/arshamalh/blogo/models"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type categoryController struct {
	basicAttributes
}

func NewCategoryController(db databases.Database, logger *zap.Logger) *categoryController {
	return &categoryController{
		basicAttributes: basicAttributes{
			db:     db,
			logger: logger,
		},
	}
}

func (cc *categoryController) CreateCategory(ctx echo.Context) error {
	var category models.Category
	if err := ctx.Bind(&category); err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusBadRequest, echo.Map{"status": "invalid request"})
	} else if cc.db.CheckCategoryExists(category.Name) {
		return ctx.JSON(http.StatusConflict, echo.Map{"status": "category already exists"})
	} else {
		_, err := cc.db.CreateCategory(&category)
		if err != nil {
			log.Gl.Error(err.Error())
			return ctx.JSON(http.StatusConflict, echo.Map{"status": "cannot create category"})
		}
		return ctx.JSON(http.StatusCreated, echo.Map{"status": "category created"})
	}
}

func (cc *categoryController) GetCategory(ctx echo.Context) error {
	category, err := cc.db.GetCategory(ctx.Param("name"))
	if err != nil || category.ID == 0 {
		return ctx.JSON(http.StatusNotFound, echo.Map{"status": "category not found"})
	}
	log.Gl.Info("Category retrieved", zap.String("category_name", category.Name))
	return ctx.JSON(http.StatusOK, category)
}

func (cc *categoryController) GetCategories(ctx echo.Context) error {
	categories, err := cc.db.GetCategories()
	if err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusNotFound, echo.Map{"status": "categories not found"})
	}

	return ctx.JSON(http.StatusOK, categories)
}
