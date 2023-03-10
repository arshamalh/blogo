package database

import "github.com/arshamalh/blogo/models"

func (gdb *gormdb) CheckCategoryExists(name string) bool {
	results := gdb.db.Take(&models.Category{}, "name = ?", name)
	return results.RowsAffected > 0
}

func (gdb *gormdb) CreateCategory(catg *models.Category) (uint, error) {
	err := gdb.db.Create(&catg).Error
	return catg.ID, err
}

func (gdb *gormdb) GetCategory(name string) (*models.Category, error) {
	catg := &models.Category{}
	err := gdb.db.Preload("Post").First(&catg, "name = ?", name).Error
	return catg, err
}

func (gdb *gormdb) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := gdb.db.Find(&categories).Error
	return categories, err
}
