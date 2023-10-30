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

func (cc *commentController) LogInfo(message string) {
	if cc.logger != nil {
		cc.logger.Info(message)
	}
}

func (cc *commentController) CreateComment(ctx echo.Context) error {
	var comment models.Comment
	userID, err := tools.ExtractUserID(ctx)
	if err != nil {
		cc.logger.Error(err.Error())
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "there is a problem with your user"})
	}

	if err := ctx.Bind(&comment); err != nil {
		cc.LogInfo("Unable to parse comment")
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "unable to parse comment"})
	}

	comment.UserID = userID
	if err := cc.db.AddComment(&comment); err != nil {
		cc.LogInfo("Unable to add comment")
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "unable to add comment"})
	}

	cc.LogInfo("Comment added successfully")
	return ctx.JSON(http.StatusCreated, echo.Map{"comment": comment})
}
