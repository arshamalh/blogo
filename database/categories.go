package database

import "github.com/arshamalh/blogo/models"

func CheckCategoryExists(name string) bool {
	results := DB.Take(&models.Category{}, "name = ?", name)
	return results.RowsAffected > 0
}

func CreateCategory(catg *models.Category) (uint, error) {
	err := DB.Create(&catg).Error
	return catg.ID, err
}

func GetCategory(name string) (*models.Category, error) {
	catg := &models.Category{}
	err := DB.Preload("Post").First(&catg, "name = ?", name).Error
	return catg, err
}

func GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := DB.Find(&categories).Error
	return categories, err
}
