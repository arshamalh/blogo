// comments.go
package controllers

import (
	"net/http"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/log"
	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/tools"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// commentController defines the controller for managing comments.
type commentController struct {
	db databases.Database
}

// NewCommentController creates a new instance of commentController.
func NewCommentController(db databases.Database) *commentController {
	return &commentController{
		db: db,
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
	// Extract user ID from the context
	userID, err := tools.ExtractUserID(ctx)
	if err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "there is a problem with your user"})
	}

	// Bind the request body to the comment model
	var comment models.Comment
	if err := ctx.Bind(&comment); err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "unable to parse comment"})
	}

	// Set the user ID for the comment
	comment.UserID = userID

	// Add the comment to the database
	if err := cc.db.AddComment(&comment); err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "unable to add comment"})
	}

	// Prepare the response
	response := map[string]interface{}{
		"comment":    comment,
		"author_id":  userID,
		"comment_id": comment.ID,
	}
	log.Gl.Info("Comment created", zap.Any("response", response))

	// Return the response with status code 201 (Created)
	return ctx.JSON(http.StatusCreated, response)
}
