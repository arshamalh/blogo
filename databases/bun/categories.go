package database

import (
	"context"

	"github.com/arshamalh/blogo/models"
)

func (bdb *bundb) CheckCategoryExists(name string) bool {
	var count int
	_, err := bdb.db.NewSelect().Model((*models.Category)(nil)).Where("name = ?", name).Count(context.Background())
	if err != nil {
		return false
	}
	return count > 0
}

func (bdb *bundb) CreateCategory(catg *models.Category) (uint, error) {
	_, err := bdb.db.NewInsert().Model(catg).Exec(context.Background())
	return catg.ID, err
}

func (bdb *bundb) GetCategory(name string) (*models.Category, error) {
	var category models.Category
	err := bdb.db.NewSelect().Model(&category).Relation("Posts").Where("name = ?", name).Scan(context.Background())
	return &category, err
}

func (bdb *bundb) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := bdb.db.NewSelect().Model(&categories).Scan(context.Background())
	return categories, err
}
