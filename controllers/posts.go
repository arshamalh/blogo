package controllers

import (
	"net/http"
	"strconv"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/log"
	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/tools"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type PostRequest struct {
	Title      string   `form:"title" json:"title" binding:"required"`
	Content    string   `form:"content" json:"content" binding:"required"`
	Categories []string `form:"categories" json:"categories"`
}

type postController struct {
	basicAttributes
}

func NewPostController(db databases.Database, logger *zap.Logger) *postController {
	return &postController{
		basicAttributes: basicAttributes{
			db:     db,
			logger: logger,
		},
	}
}

func (pc *postController) LogInfo(message string) {
	if log.Gl != nil {
		log.Gl.Info(message)
	}
}

func (pc *postController) CreatePost(ctx echo.Context) error {
	var post PostRequest
	userID, _ := tools.ExtractUserID(ctx)
	if err := ctx.Bind(&post); err != nil {
		log.Gl.Info("Unable to parse post")
		return ctx.JSON(500, echo.Map{"error": err.Error()})
	}

	categories := []models.Category{}
	for _, category := range post.Categories {
		categories = append(categories, models.Category{Name: category})
	}

	newPost := models.Post{
		Title:      post.Title,
		Content:    post.Content,
		AuthorID:   userID,
		Categories: categories,
	}
	postID, err := pc.db.CreatePost(&newPost)
	if err != nil {
		log.Gl.Error("An error occurred", zap.Error(err))
		return ctx.JSON(http.StatusConflict, echo.Map{"message": "cannot make the post"})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{
		"message": "Post created successfully",
		"post_id": postID,
	})
}

func (pc *postController) DeletePost(ctx echo.Context) error {
	userID, _ := tools.ExtractUserID(ctx)
	postID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	// If the user doesn't have the permission to delete posts, check if it's their own post.
	if !tools.ExtractPermissable(ctx) {
		post, _ := pc.db.GetPost(uint(postID))
		if post.AuthorID != userID {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{
				"message": "you don't have enough permissions",
			})
		}
	}

	if err := pc.db.DeletePost(uint(postID)); err != nil {
		log.Gl.Error("An error occurred", zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "Post deleted"})
}

func (pc *postController) UpdatePost(ctx echo.Context) error {
	var post PostRequest
	if err := ctx.Bind(&post); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"status": "invalid request"})
	}

	// Make a new updated post
	categories := []models.Category{}
	for _, category := range post.Categories {
		categories = append(categories, models.Category{Name: category})
	}
	newPost := models.Post{
		Title:      post.Title,
		Content:    post.Content,
		Categories: categories,
	}
	postID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	newPost.ID = uint(postID)

	// Update post
	if err := pc.db.UpdatePost(&newPost); err != nil {
		return ctx.JSON(500, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(200, echo.Map{"post": newPost})
}

func (pc *postController) GetPost(ctx echo.Context) error {
	postID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	post, _ := pc.db.GetPost(uint(postID))
	return ctx.JSON(200, echo.Map{"post": post})
}

func (pc *postController) GetPosts(ctx echo.Context) error {
	posts, _ := pc.db.GetPosts()
	return ctx.JSON(200, echo.Map{"posts": posts})
}
