package controllers

import (
	"strconv"

	"github.com/arshamalh/blogo/database"
	"github.com/arshamalh/blogo/models"
	"github.com/gin-gonic/gin"
)

func CreateRole(ctx *gin.Context) {
	var role models.Role
	if err := ctx.BindJSON(&role); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if err := database.CreateRole(&role); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"role": role})
}

func UpdateRole(ctx *gin.Context) {
	var role models.Role
	if err := ctx.BindJSON(&role); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if err := database.UpdateRole(&role); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"role": role})
}

func DeleteRole(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := database.DeleteRole(uint(id)); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Role deleted"})
}

func GetRole(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	role, err := database.GetRole(uint(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"role": role})
}

func GetRoles(ctx *gin.Context) {
	roles, err := database.GetRoles()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"roles": roles})
}
