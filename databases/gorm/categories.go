package database

import (
	"log"

	"github.com/arshamalh/blogo/models"
)

func (gdb *gormdb) CheckCategoryExists(name string) bool {
	results := gdb.db.Take(&models.Category{}, "name = ?", name)
	exists := results.RowsAffected > 0

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
	catg := &models.Category{}
	err := gdb.db.Preload("Post").First(&catg, "name = ?", name).Error

	if err == nil {
		log.Printf("Retrieved category with name '%s'", name)
	} else {
		log.Printf("Failed to retrieve category with name '%s'. Error: %v", name, err)
	}

	return catg, err
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
