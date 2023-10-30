package controllers

import (
	"github.com/arshamalh/blogo/databases"
	"go.uber.org/zap"
)

type basicAttributes struct {
	db     databases.Database
	logger *zap.Logger
}

func (ba *basicAttributes) LogInfo(message string) {
	if ba.logger != nil {
		ba.logger.Info(message)
	}
}
