package controllers

import (
	"net/http"
	"strconv"

	"github.com/arshamalh/blogo/databases"
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

func (pc *postController) CreatePost(ctx echo.Context) error {
	var post PostRequest
	user_id, _ := tools.ExtractUserID(ctx)
	if err := ctx.Bind(&post); err != nil {
		return ctx.JSON(500, echo.Map{"error": err.Error()})
	}

	catgs := []models.Category{}
	for _, catg := range post.Categories {
		catgs = append(catgs, models.Category{Name: catg})
	}

	new_post := models.Post{
		Title:      post.Title,
		Content:    post.Content,
		AuthorID:   user_id,
		Categories: catgs,
	}
	post_id, err := pc.db.CreatePost(&new_post)
	if err != nil {
		pc.logger.Error(err.Error())
		return ctx.JSON(http.StatusConflict, echo.Map{"message": "cannot make the post"})
	}
	return ctx.JSON(http.StatusCreated, echo.Map{
		"message": "Post created successfully",
		"post_id": post_id,
	})
}

func (pc *postController) DeletePost(ctx echo.Context) error {
	user_id, _ := tools.ExtractUserID(ctx)
	post_id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	// If user doesn't have the permission to delete posts, check if it's its own post.
	if !tools.ExtractPermissable(ctx) {
		post, _ := pc.db.GetPost(uint(post_id))
		if post.AuthorID != user_id {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{
				"message": "you don't have enough permissions",
			})
		}
	}

	// Delete post
	if err := pc.db.DeletePost(uint(post_id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})

	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "Post deleted"})
}

func (pc *postController) UpdatePost(ctx echo.Context) error {
	var post PostRequest
	if err := ctx.Bind(&post); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"status": "invalid request"})

	}

	// Make new updated post
	catgs := []models.Category{}
	for _, catg := range post.Categories {
		catgs = append(catgs, models.Category{Name: catg})
	}
	new_post := models.Post{
		Title:      post.Title,
		Content:    post.Content,
		Categories: catgs,
	}
	post_id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	new_post.ID = uint(post_id)

	// Update post
	if err := pc.db.UpdatePost(&new_post); err != nil {
		return ctx.JSON(500, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(200, echo.Map{"post": new_post})
}

func (pc *postController) GetPost(ctx echo.Context) error {
	post_id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	post, _ := pc.db.GetPost(uint(post_id))
	return ctx.JSON(200, echo.Map{
		"post": post,
	})
}

func (pc *postController) GetPosts(ctx echo.Context) error {
	posts, _ := pc.db.GetPosts()
	return ctx.JSON(200, echo.Map{
		"posts": posts,
	})
}
