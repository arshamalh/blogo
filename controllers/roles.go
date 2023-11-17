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

// CreateRole creates a new user role.
// swagger:route POST /roles CreateRole
// Creates a new user role.
// responses:
//
//	201: map[string]interface{} "Role created"
//	400: map[string]interface{} "Invalid request"
//	500: map[string]interface{} "Internal server error"
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

// UpdateRole updates an existing user role.
// swagger:route PUT /roles UpdateRole
// Updates an existing user role.
// responses:
//
//	200: map[string]interface{} "Role updated"
//	400: map[string]interface{} "Invalid request"
//	500: map[string]interface{} "Internal server error"
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

// DeleteRole deletes an existing user role.
// swagger:route DELETE /roles/{id} DeleteRole
// Deletes an existing user role.
// parameters:
//   - name: id
//     in: path
//     description: The ID of the role to delete.
//     required: true
//     type: integer
//
// responses:
//
//	200: map[string]interface{} "Role deleted"
//	500: map[string]interface{} "Internal server error"
func (rc *roleController) DeleteRole(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := rc.db.DeleteRole(uint(id)); err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	log.Gl.Info("Role deleted", zap.Uint64("role_id", id))
	return ctx.JSON(http.StatusOK, echo.Map{"message": "Role deleted"})
}

// GetRole retrieves a user role by ID.
// swagger:route GET /roles/{id} GetRole
// Retrieves a user role by ID.
// parameters:
//   - name: id
//     in: path
//     description: The ID of the role to retrieve.
//     required: true
//     type: integer
//
// responses:
//
//	200: map[string]interface{} "Role retrieved"
//	404: map[string]interface{} "Role not found"
//	500: map[string]interface{} "Internal server error"
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

// GetRoles retrieves all user roles.
// swagger:route GET /roles GetRoles
// Retrieves all user roles.
// responses:
//
//	200: map[string]interface{} "Roles retrieved"
//	500: map[string]interface{} "Internal server error"
func (rc *roleController) GetRoles(ctx echo.Context) error {
	roles, err := rc.db.GetRoles()
	if err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	log.Gl.Info("Roles retrieved", zap.Int("total_roles", len(roles)))
	return ctx.JSON(http.StatusOK, echo.Map{"roles": roles})
}
