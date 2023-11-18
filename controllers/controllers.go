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

// @title Blogo API server
// @version 1.0
// @description API for managing categories and comments in Blogo application.
// @contact.email arshamalh.github.io/
// @contact.url https://blogo.com/contact
// @BasePath /api/v1
