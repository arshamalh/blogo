package controllers

import (
	"net/http"
	"strconv"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/models"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type roleController struct {
	db databases.Database
}

func NewRoleController(db databases.Database) *roleController {
	return &roleController{
		db: db,
	}
}

func (rc *roleController) CreateRole(ctx echo.Context) error {
	var role models.Role
	if err := ctx.Bind(&role); err != nil {
		return ctx.JSON(http.StatusBadRequest, gin.H{"status": "invalid request"})
	}
	if err := rc.db.CreateRole(&role); err != nil {
		return ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, gin.H{"role": role})
}

func (rc *roleController) UpdateRole(ctx echo.Context) error {
	var role models.Role
	if err := ctx.Bind(&role); err != nil {
		return ctx.JSON(http.StatusBadRequest, gin.H{"status": "invalid request"})

	}
	if err := rc.db.UpdateRole(&role); err != nil {
		return ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}
	return ctx.JSON(http.StatusOK, gin.H{"role": role})
}

func (rc *roleController) DeleteRole(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := rc.db.DeleteRole(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}
	return ctx.JSON(http.StatusOK, gin.H{"message": "Role deleted"})
}

func (rc *roleController) GetRole(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	role, err := rc.db.GetRole(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}
	return ctx.JSON(http.StatusOK, gin.H{"role": role})
}

func (rc *roleController) GetRoles(ctx echo.Context) error {
	roles, err := rc.db.GetRoles()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}
	return ctx.JSON(http.StatusOK, gin.H{"roles": roles})
}
