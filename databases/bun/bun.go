package database

import (
	"database/sql"

	"github.com/arshamalh/blogo/log"
	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/models/permissions"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type bundb struct {
	db *bun.DB
}

func Connect(dsn string) (*bundb, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	err := db.Ping()
	if err != nil {
		return nil, err
	}

	log.Gl.Info("Database connection successfully opened")

	return &bundb{db: db}, nil
}

// Add some basic roles manually
func (bundb *bundb) AddBasicRoles() {
	bundb.CreateRole(&models.Role{Name: "superadmin", Permissions: permissions.Compress([]permissions.Permission{permissions.FullAccess})})
	bundb.CreateRole(&models.Role{Name: "moderator", Permissions: permissions.Compress([]permissions.Permission{permissions.FullContents})})
	bundb.CreateRole(&models.Role{Name: "author", Permissions: permissions.Compress([]permissions.Permission{permissions.CreatePost, permissions.FullContents})})
}
