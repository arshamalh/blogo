// controllers.go
package controllers

import (
	"github.com/arshamalh/blogo/databases"
	"go.uber.org/zap"
)

// basicAttributes defines the basic attributes shared by other controllers.
type basicAttributes struct {
	db     databases.Database
	logger *zap.Logger
	Gl     *zap.Logger
}

// NewBasicAttributesController creates a new instance of basicAttributesController.
func NewBasicAttributesController(db databases.Database, logger *zap.Logger) *basicAttributes {
	return &basicAttributes{
		db:     db,
		logger: logger,
		Gl:     logger,
	}
}
