package controllers

import (
	"net/http"
	"strconv"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/models"
	"github.com/gin-gonic/gin"
)

type roleController struct {
	db databases.Database
}

func NewRoleController(db databases.Database) *roleController {
	return &roleController{
		db: db,
	}
}

func (rc *roleController) CreateRole(ctx *gin.Context) {
	var role models.Role
	if err := ctx.BindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "invalid request"})
		return
	}
	if err := rc.db.CreateRole(&role); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"role": role})
}

func (rc *roleController) UpdateRole(ctx *gin.Context) {
	var role models.Role
	if err := ctx.BindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "invalid request"})
		return
	}
	if err := rc.db.UpdateRole(&role); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"role": role})
}

func (rc *roleController) DeleteRole(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := rc.db.DeleteRole(uint(id)); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Role deleted"})
}

func (rc *roleController) GetRole(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	role, err := rc.db.GetRole(uint(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"role": role})
}

func (rc *roleController) GetRoles(ctx *gin.Context) {
	roles, err := rc.db.GetRoles()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"roles": roles})
}
