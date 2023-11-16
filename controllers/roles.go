package controllers

import (
	"net/http"
	"strconv"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/log"
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
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusBadRequest, echo.Map{"status": "invalid request"})
	}
	if err := rc.db.CreateRole(&role); err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	log.Gl.Info("Role created", zap.String("role_name", role.Name))
	return ctx.JSON(http.StatusCreated, echo.Map{"role": role})
}

func (rc *roleController) UpdateRole(ctx echo.Context) error {
	var role models.Role
	if err := ctx.Bind(&role); err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusBadRequest, echo.Map{"status": "invalid request"})
	}
	if err := rc.db.UpdateRole(&role); err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	log.Gl.Info("Role updated", zap.String("role_name", role.Name))
	return ctx.JSON(http.StatusOK, echo.Map{"role": role})
}

func (rc *roleController) DeleteRole(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := rc.db.DeleteRole(uint(id)); err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	log.Gl.Info("Role deleted", zap.Uint64("role_id", id))
	return ctx.JSON(http.StatusOK, echo.Map{"message": "Role deleted"})
}

func (rc *roleController) GetRole(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	role, err := rc.db.GetRole(uint(id))
	if err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	log.Gl.Info("Role retrieved", zap.Uint64("role_id", id))
	return ctx.JSON(http.StatusOK, echo.Map{"role": role})
}

func (rc *roleController) GetRoles(ctx echo.Context) error {
	roles, err := rc.db.GetRoles()
	if err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	log.Gl.Info("Roles retrieved", zap.Int("total_roles", len(roles)))
	return ctx.JSON(http.StatusOK, echo.Map{"roles": roles})
}
