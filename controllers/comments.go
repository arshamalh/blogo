// comments.go
package controllers

import (
	"net/http"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/tools"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// commentController defines the controller for managing comments.
type commentController struct {
	basicAttributes
}

// NewCommentController creates a new instance of commentController.
func NewCommentController(db databases.Database, logger *zap.Logger) *commentController {
	return &commentController{
		basicAttributes: basicAttributes{
			db:     db,
			logger: logger,
		},
	}
}

// @Summary Create a new comment
// @Description Create a new comment
// @ID create-comment
// @Accept json
// @Produce json
// @Param comment body models.Comment true "Comment object"
// @Success 201 {object} map[string]any
// @Failure 400 {object} map[string]any
// @Failure 401 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Router /comments [post]
func (cc *commentController) CreateComment(ctx echo.Context) error {
	var comment models.Comment
	user_id, err := tools.ExtractUserID(ctx)
	if err != nil {
		cc.logger.Error(err.Error())
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "there is problem with your user"})
	}
	if err := ctx.Bind(&comment); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "unable to parse comment"})
	}
	comment.UserID = user_id
	if err := cc.db.AddComment(&comment); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "unable to add comment"})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{"comment": comment})
}
