// category_controller.go
package controllers

import (
	"net/http"

	"github.com/arshamalh/blogo/databases"
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

// @Summary Create a new category
// @Description Create a new category
// @ID create-category
// @Accept json
// @Produce json
// @Param category body models.Category true "Category object"
// @Success 201 {object} map[string]any
// @Failure 400 {object} map[string]any
// @Router /categories [post]
func (cc *categoryController) CreateCategory(ctx echo.Context) error {
	var category models.Category
	if err := ctx.Bind(&category); err != nil {
		cc.logger.Error(err.Error())
		return ctx.JSON(http.StatusBadRequest, echo.Map{"status": "invalid request"})
	} else if cc.db.CheckCategoryExists(category.Name) {
		return ctx.JSON(http.StatusConflict, echo.Map{"status": "category already exists"})
	} else {
		_, err := cc.db.CreateCategory(&category)
		if err != nil {
			cc.logger.Error(err.Error())
			ctx.JSON(http.StatusConflict, echo.Map{"status": "cannot create category"})
		}
		return ctx.JSON(http.StatusCreated, echo.Map{"status": "category created"})
	}
}

// @Summary Get a category by name
// @Description Get category details by name
// @ID get-category
// @Param name path string true "Category Name"
// @Produce json
// @Success 200 {object} models.Category
// @Failure 404 {object} map[string]any
// @Router /categories/{name} [get]
func (cc *categoryController) GetCategory(ctx echo.Context) error {
	category, err := cc.db.GetCategory(ctx.Param("name"))
	if err != nil || category.ID == 0 {
		cc.logger.Error(err.Error())
		return ctx.JSON(http.StatusNotFound, echo.Map{"status": "category not found"})
	}
	return ctx.JSON(http.StatusOK, category)
}

// @Summary Get all categories
// @Description Get a list of all categories
// @ID get-all-categories
// @Produce json
// @Success 200 {array} models.Category
// @Failure 404 {object} map[string]any
// @Router /categories [get]
func (cc *categoryController) GetCategories(ctx echo.Context) error {
	categories, err := cc.db.GetCategories()
	if err != nil {
		cc.logger.Error(err.Error())
		return ctx.JSON(http.StatusNotFound, echo.Map{"status": "categories not found"})
	}
	return ctx.JSON(http.StatusOK, categories)
}
