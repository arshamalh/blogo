package middlewares

import (
	"github.com/arshamalh/blogo/database"
	"github.com/arshamalh/blogo/models/permissions"
	"github.com/arshamalh/blogo/tools"
	"github.com/gin-gonic/gin"
)

func CheckPermissions(permission permissions.Permission) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user_id, err := tools.ExtractUserID(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		for _, perm := range database.GetUserPermissions(user_id) {
			if perm == permission ||
				perm == permissions.FullAccess ||
				(perm == permissions.FullContents && permission > permissions.FullContents) {
				ctx.Next()
				return
			}
		}
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "you don't have enough permissions",
		})
	}
}
