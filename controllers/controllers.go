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

func (ba *basicAttributes) LogInfo(message string) {
	if ba.Gl != nil {
		ba.Gl.Info(message)
	}
}

func (ba *basicAttributes) LogError(message string) {
	if ba.Gl != nil {
		ba.Gl.Error(message)
	}
}
