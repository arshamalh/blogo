// roles.go
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

// roleController handles HTTP requests related to user roles.
type roleController struct {
	basicAttributes
}

// NewRoleController creates a new instance of roleController.
func NewRoleController(db databases.Database, logger *zap.Logger) *roleController {
	return &roleController{
		basicAttributes: basicAttributes{
			db:     db,
			logger: logger,
		},
	}
}

// @Summary Create a new user role
// @Description Create a new user role
// @ID create-role
// @Accept json
// @Produce json
// @Param role body models.Role true "Role object"
// @Success 201 {object} map[string]interface{} "Role created"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /roles [post]
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

// @Summary Update an existing user role
// @Description Update an existing user role
// @ID update-role
// @Accept json
// @Produce json
// @Param role body models.Role true "Role object"
// @Success 200 {object} map[string]interface{} "Role updated"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /roles [put]
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

// @Summary Delete an existing user role
// @Description Delete an existing user role
// @ID delete-role
// @Param id path int true "Role ID"
// @Success 200 {object} map[string]interface{} "Role deleted"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /roles/{id} [delete]
func (rc *roleController) DeleteRole(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := rc.db.DeleteRole(uint(id)); err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	log.Gl.Info("Role deleted", zap.Uint64("role_id", id))
	return ctx.JSON(http.StatusOK, echo.Map{"message": "Role deleted"})
}

// @Summary Retrieve a user role by ID
// @Description Retrieve a user role by ID
// @ID get-role
// @Param id path int true "Role ID"
// @Produce json
// @Success 200 {object} map[string]interface{} "Role retrieved"
// @Failure 404 {object} map[string]interface{} "Role not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /roles/{id} [get]
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

// @Summary Retrieve all user roles
// @Description Retrieve all user roles
// @ID get-roles
// @Produce json
// @Success 200 {object} map[string]interface{} "Roles retrieved"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /roles [get]
func (rc *roleController) GetRoles(ctx echo.Context) error {
	roles, err := rc.db.GetRoles()
	if err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	log.Gl.Info("Roles retrieved", zap.Int("total_roles", len(roles)))
	return ctx.JSON(http.StatusOK, echo.Map{"roles": roles})
}
