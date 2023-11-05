package database

import (
	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/models/permissions"
	"go.uber.org/zap"

	"github.com/arshamalh/blogo/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormdb struct {
	db *gorm.DB
}

func NewGormDB(db *gorm.DB) *gormdb {
	return &gormdb{
		db: db,
	}
}

func Connect(dsn string) (*gormdb, error) {
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Gl.Error("Failed to connect to database: ", zap.Error(err))
		return nil, err
	}

	log.Gl.Info("Database connection successfully opened")

	// Auto migration
	if err := DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Category{}, &models.Role{}, &models.Comment{}); err != nil {
		log.Gl.Info("Failed to migrate the database: %v", zap.Error(err))
		return nil, err
	}
	log.Gl.Info("Database Migrated")
	gdb := NewGormDB(DB)
	return gdb, nil
}

func (gdb *gormdb) AddBasicRoles() {
	// Creating superadmin role
	superadminRole := &models.Role{Name: "superadmin", Permissions: permissions.Compress([]permissions.Permission{permissions.FullAccess})}
	if err := gdb.CreateRole(superadminRole); err != nil {
		log.Gl.Info("Failed to create role 'superadmin': %v", zap.Error(err))
	}

	// Creating moderator role
	moderatorRole := &models.Role{Name: "moderator", Permissions: permissions.Compress([]permissions.Permission{permissions.FullContents})}
	if err := gdb.CreateRole(moderatorRole); err != nil {
		log.Gl.Info("Failed to create role 'moderator': %v", zap.Error(err))
	}

	// Creating author role
	authorRole := &models.Role{Name: "author", Permissions: permissions.Compress([]permissions.Permission{permissions.CreatePost, permissions.FullContents})}
	if err := gdb.CreateRole(authorRole); err != nil {
		log.Gl.Info("Failed to create role 'author': %v", zap.Error(err))
	}
}

func (gdb *gormdb) CreateRole(role *models.Role) error {
	err := gdb.db.Create(role).Error
	if err == nil {
		log.Gl.Info("Role '" + role.Name + "' created successfully")
	} else {
		log.Gl.Info("Failed to create role '"+role.Name+"': ", zap.Error(err))
	}
	return err
}
