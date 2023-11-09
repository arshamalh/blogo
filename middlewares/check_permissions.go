package middlewares

import (
	"fmt"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/log"
	"github.com/arshamalh/blogo/models/permissions"
	"github.com/arshamalh/blogo/tools"
	"github.com/labstack/echo/v4"
)

// CheckPermissions is a middleware that checks if the user has the required permissions.
func CheckPermissions(db databases.Database, permission permissions.Permission) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			userID, _ := tools.ExtractUserID(ctx)
			hasPermission := HavePermissions(db, userID, permission)

			// Log information
			if hasPermission {
				log.Gl.Info(fmt.Sprintf("User with ID %d has the required permission: %s", userID, permission))
			}

			ctx.Set("permissable", hasPermission)
			return next(ctx)
		}
	}
}

// HavePermissions checks if the user has the required permissions, this can be used in other parts of the application.
func HavePermissions(db databases.Database, userID uint, permission permissions.Permission) bool {
	for _, perm := range db.GetUserPermissions(userID) {
		if perm == permission ||
			perm == permissions.FullAccess ||
			(perm == permissions.FullContents && permission > permissions.FullContents) {
			return true
		}
	}
	return false
}
