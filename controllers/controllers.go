package controllers

import (
	"github.com/arshamalh/blogo/databases"
	"go.uber.org/zap"
)

type basicAttributes struct {
	db     databases.Database
	logger *zap.Logger
	Gl     *zap.Logger
}
