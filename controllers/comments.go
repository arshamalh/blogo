package controllers

import (
	"net/http"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/log"
	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/tools"
	"github.com/labstack/echo/v4"
)

type commentController struct {
	db databases.Database
}

func NewCommentController(db databases.Database) *commentController {
	return &commentController{
		db: db,
	}
}

func (cc *commentController) CreateComment(ctx echo.Context) error {
	var comment models.Comment
	userID, err := tools.ExtractUserID(ctx)
	if err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "there is a problem with your user"})
	}

	if err := ctx.Bind(&comment); err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "unable to parse comment"})
	}

	comment.UserID = userID

	if err := cc.db.AddComment(&comment); err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "unable to add comment"})
	}

	response := map[string]interface{}{
		"comment":    comment,
		"author_id":  userID,
		"comment_id": comment.ID,
	}

	return ctx.JSON(http.StatusCreated, response)
}
