package models

import (
	"github.com/arshamalh/blogo/models/permissions"
	"gorm.io/gorm"
)

// Any user has a role, and any role has some permissions and accessories in the future (TODO) (logo, cool name and more)
type Role struct {
	gorm.Model
	Name        string                   `form:"name" json:"name" binding:"required"`
	Permissions []permissions.Permission `form:"premissions" json:"premissions"`
}
