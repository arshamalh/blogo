package database

import (
	"errors"
	"log"

	"github.com/arshamalh/blogo/models"
	"gorm.io/gorm"
)

func (gdb *gormdb) CheckCategoryExists(name string) bool {
	var category models.Category
	err := gdb.db.Where("name = ?", name).First(&category).Error
	exists := !errors.Is(err, gorm.ErrRecordNotFound)

	if exists {
		log.Printf("Category with name '%s' exists.", name)
	} else {
		log.Printf("Category with name '%s' does not exist.", name)
	}

	return exists
}

func (gdb *gormdb) CreateCategory(catg *models.Category) (uint, error) {
	err := gdb.db.Create(&catg).Error

	if err == nil {
		log.Printf("Created category with ID: %d", catg.ID)
	} else {
		log.Printf("Failed to create category. Error: %v", err)
	}

	return catg.ID, err
}

func (gdb *gormdb) GetCategory(name string) (*models.Category, error) {
	var category models.Category
	err := gdb.db.Where("name = ?", name).Preload("Post").First(&category).Error

	if err == nil {
		log.Printf("Retrieved category with name '%s'", name)
	} else {
		log.Printf("Failed to retrieve category with name '%s'. Error: %v", name, err)
	}

	return &category, err
}

func (gdb *gormdb) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := gdb.db.Find(&categories).Error

	if err == nil {
		log.Println("Retrieved all categories.")
	} else {
		log.Printf("Failed to retrieve categories. Error: %v", err)
	}

	return categories, err
}
