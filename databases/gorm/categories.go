package database

import (
	"github.com/arshamalh/blogo/log"
	"github.com/arshamalh/blogo/models"
)

func (gdb *gormdb) CheckCategoryExists(name string) bool {
	var category models.Category
	results := gdb.db.Where("name = ?", name).First(&category)
	if exists := results.RowsAffected > 0; exists {
		log.Gl.Info("Category with name '" + name + "' exists.")
		return exists // return right here so if exists, we don't go further
	}
	log.Gl.Info("Category with name '" + name + "' does not exist.")
	return false
}

func (gdb *gormdb) CreateCategory(catg *models.Category) (uint, error) {
	err := gdb.db.Create(&catg).Error

	return catg.ID, err
}

func (gdb *gormdb) GetCategory(name string) (*models.Category, error) {
	var category models.Category
	err := gdb.db.Where("name = ?", name).Preload("Post").First(&category).Error

	if err == nil {
		log.Gl.Info("Retrieved category with name '" + name + "'")
	} else {
		log.Gl.Error("Failed to retrieve category with name '" + name + "'. Error: " + err.Error())
	}

	return &category, err
}

func (gdb *gormdb) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := gdb.db.Find(&categories).Error

	if err == nil {
		log.Gl.Info("Retrieved all categories.")
	} else {
		log.Gl.Error("Failed to retrieve categories. Error: " + err.Error())
	}

	return categories, err
}
