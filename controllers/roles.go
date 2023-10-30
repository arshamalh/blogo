package controllers

import (
	"net/http"
	"strconv"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/models"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type roleController struct {
	basicAttributes
}

func NewRoleController(db databases.Database, logger *zap.Logger) *roleController {
	return &roleController{
		basicAttributes: basicAttributes{
			db:     db,
			logger: logger,
		},
	}
}

func (rc *roleController) CreateRole(ctx echo.Context) error {
	var role models.Role
	if err := ctx.Bind(&role); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"status": "invalid request"})
	}
	if err := rc.db.CreateRole(&role); err != nil {
		rc.LogInfo("Failed to create role: " + err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	rc.LogInfo("Role created: " + role.Name)
	return ctx.JSON(http.StatusCreated, echo.Map{"role": role})
}

func (rc *roleController) UpdateRole(ctx echo.Context) error {
	var role models.Role
	if err := ctx.Bind(&role); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"status": "invalid request"})

	}
	if err := rc.db.UpdateRole(&role); err != nil {
		rc.LogInfo("Failed to update role: " + err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})

	}
	rc.LogInfo("Role updated: " + role.Name)
	return ctx.JSON(http.StatusOK, echo.Map{"role": role})
}

func (rc *roleController) DeleteRole(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := rc.db.DeleteRole(uint(id)); err != nil {
		rc.LogInfo("Failed to delete role: " + err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})

	}
	rc.LogInfo("Role deleted: ID " + strconv.FormatUint(id, 10))
	return ctx.JSON(http.StatusOK, echo.Map{"message": "Role deleted"})
}

func (rc *roleController) GetRole(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	role, err := rc.db.GetRole(uint(id))
	if err != nil {
		rc.LogInfo("Failed to get role: " + err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})

	}
	rc.LogInfo("Role retrieved: ID " + strconv.FormatUint(id, 10))
	return ctx.JSON(http.StatusOK, echo.Map{"role": role})
}

func (rc *roleController) GetRoles(ctx echo.Context) error {
	roles, err := rc.db.GetRoles()
	if err != nil {
		rc.LogInfo("Failed to get roles: " + err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})

	}
	rc.LogInfo("Roles retrieved: Total " + strconv.Itoa(len(roles)))
	return ctx.JSON(http.StatusOK, echo.Map{"roles": roles})
}

func (rc *roleController) LogInfo(message string) {
	if rc.logger != nil {
		rc.logger.Info(message)
	}
}
