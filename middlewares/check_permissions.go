package middlewares

import (
	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/models/permissions"
	"github.com/arshamalh/blogo/tools"
	"github.com/gin-gonic/gin"
)

// CheckPermissions is a middleware that checks if the user has the required permissions.
func CheckPermissions(db databases.Database, permission permissions.Permission) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user_id, _ := tools.ExtractUserID(ctx)
		ctx.Set("permissable", HavePermissions(db, user_id, permission))
		ctx.Next()
	}
}

// HavePermissions checks if the user has the required permissions, this can be used in other parts of the application.
func HavePermissions(db databases.Database, user_id uint, permission permissions.Permission) bool {
	for _, perm := range db.GetUserPermissions(user_id) {
		if perm == permission ||
			perm == permissions.FullAccess ||
			(perm == permissions.FullContents && permission > permissions.FullContents) {
			return true
		}
	}
	return false
}
