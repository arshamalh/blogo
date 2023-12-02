package database

import (
	"context"

	"github.com/arshamalh/blogo/models"
)

func (bdb *bundb) CreateRole(role *models.Role) error {
	_, err := bdb.db.NewInsert().Model(role).Exec(context.Background())
	return err
}

func (bdb *bundb) UpdateRole(role *models.Role) error {
	_, err := bdb.db.NewUpdate().Model(role).Exec(context.Background())
	return err
}

func (bdb *bundb) DeleteRole(id uint) error {
	_, err := bdb.db.NewDelete().Model((*models.Role)(nil)).Where("id = ?", id).Exec(context.Background())
	return err
}

func (bdb *bundb) GetRole(id uint) (models.Role, error) {
	var role models.Role
	err := bdb.db.NewSelect().Model(&role).Where("id = ?", id).Scan(context.Background())
	return role, err
}

func (bdb *bundb) GetRoles() ([]models.Role, error) {
	var roles []models.Role
	err := bdb.db.NewSelect().Model(&roles).Scan(context.Background())
	return roles, err
}
