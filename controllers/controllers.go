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
}
