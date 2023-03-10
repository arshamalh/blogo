package controllers

import (
	"net/http"
	"strconv"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/tools"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type PostRequest struct {
	Title      string   `form:"title" json:"title" binding:"required"`
	Content    string   `form:"content" json:"content" binding:"required"`
	Categories []string `form:"categories" json:"categories"`
}

type postController struct {
	db databases.Database
}

func NewPostController(db databases.Database) *postController {
	return &postController{
		db: db,
	}
}

func (pc *postController) CreatePost(ctx echo.Context) error {
	var post PostRequest
	user_id, _ := tools.ExtractUserID(ctx)
	if err := ctx.Bind(&post); err != nil {
		return ctx.JSON(500, gin.H{"error": err.Error()})
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
	post_id, _ := pc.db.CreatePost(&new_post)
	return ctx.JSON(200, gin.H{
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
			return ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "you don't have enough permissions",
			})
		}
	}

	// Delete post
	if err := pc.db.DeletePost(uint(post_id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}
	return ctx.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

func (pc *postController) UpdatePost(ctx echo.Context) error {
	var post PostRequest
	if err := ctx.Bind(&post); err != nil {
		return ctx.JSON(http.StatusBadRequest, gin.H{"status": "invalid request"})

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
		return ctx.JSON(500, gin.H{"error": err.Error()})
	}
	return ctx.JSON(200, gin.H{"post": new_post})
}

func (pc *postController) GetPost(ctx echo.Context) error {
	post_id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	post, _ := pc.db.GetPost(uint(post_id))
	return ctx.JSON(200, gin.H{
		"post": post,
	})
}

func (pc *postController) GetPosts(ctx echo.Context) error {
	posts, _ := pc.db.GetPosts()
	return ctx.JSON(200, gin.H{
		"posts": posts,
	})
}
