package database

import "github.com/arshamalh/blogo/models"

func (gdb *gormdb) CreateRole(role *models.Role) error {
	err := gdb.db.Create(&role).Error
	return err
}

func (gdb *gormdb) UpdateRole(role *models.Role) error {
	err := gdb.db.Save(&role).Error
	return err
}

func (gdb *gormdb) DeleteRole(id uint) error {
	err := gdb.db.Delete(&models.Role{}, "id = ?", id).Error
	return err
}

func (gdb *gormdb) GetRole(id uint) (models.Role, error) {
	var role models.Role
	err := gdb.db.First(&role, id).Error
	return role, err
}

func (gdb *gormdb) GetRoles() ([]models.Role, error) {
	var roles []models.Role
	err := gdb.db.Find(&roles).Error
	return roles, err
}
