package database

import (
	"fmt"

	"github.com/arshamalh/blogo/log"
	"github.com/arshamalh/blogo/models"
)

func (gdb *gormdb) CreateNewRole(role *models.Role) error {
	err := gdb.db.Create(role).Error

	if err == nil {
		log.Gl.Info(fmt.Sprintf("Role created with ID: %d", role.ID))
	}
	return err
}

func (gdb *gormdb) UpdateRole(role *models.Role) error {
	err := gdb.db.Save(role).Error

	if err == nil {
		log.Gl.Info(fmt.Sprintf("Role updated with ID: %d", role.ID))
	}

	return err
}

func (gdb *gormdb) DeleteRole(id uint) error {
	err := gdb.db.Delete(&models.Role{}, id).Error

	if err == nil {
		log.Gl.Info(fmt.Sprintf("Role with ID %d deleted", id))
	}

	return err
}

func (gdb *gormdb) GetRole(id uint) (models.Role, error) {
	var role models.Role
	err := gdb.db.First(&role, id).Error

	if err != nil {
		log.Gl.Error(fmt.Sprintf("Error: %d: %v", id, err))
	}

	return role, err
}

func (gdb *gormdb) GetRoles() ([]models.Role, error) {
	var roles []models.Role
	err := gdb.db.Find(&roles).Error

	if err != nil {
		log.Gl.Error(fmt.Sprintf("Error: %v", err))
	}

	return roles, err
}
