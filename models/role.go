package models

import (
	"gorm.io/gorm"
)

// Any user has a role, and any role has some permissions and accessories in the future (TODO) (logo, cool name and more)
type Role struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex" form:"name" json:"name" binding:"required"`
	Permissions string `form:"premissions" json:"premissions"`
}
