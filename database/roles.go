package database

import "github.com/arshamalh/blogo/models"

func CreateRole(role *models.Role) error {
	err := DB.Create(&role).Error
	return err
}

func UpdateRole(role *models.Role) error {
	err := DB.Save(&role).Error
	return err
}

func DeleteRole(id uint) error {
	err := DB.Delete(&models.Role{}, "id = ?", id).Error
	return err
}

func GetRole(id uint) (models.Role, error) {
	var role models.Role
	err := DB.First(&role, id).Error
	return role, err
}

func GetRoles() ([]models.Role, error) {
	var roles []models.Role
	err := DB.Find(&roles).Error
	return roles, err
}
