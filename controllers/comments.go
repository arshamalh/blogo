package controllers

import (
	"net/http"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/tools"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type commentController struct {
	basicAttributes
}

func NewCommentController(db databases.Database, logger *zap.Logger) *commentController {
	return &commentController{
		basicAttributes: basicAttributes{
			db:     db,
			logger: logger,
		},
	}
}

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
