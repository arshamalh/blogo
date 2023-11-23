// posts.go
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

// PostRequest represents the request structure for creating a post.
// swagger:parameters CreatePost
type PostRequest struct {
	Title      string   `form:"title" json:"title" binding:"required"`
	Content    string   `form:"content" json:"content" binding:"required"`
	Categories []string `form:"categories" json:"categories"`
}

// postController handles HTTP requests related to blog posts.
type postController struct {
	basicAttributes
}

// NewPostController creates a new instance of postController.
func NewPostController(db databases.Database, logger *zap.Logger) *postController {
	return &postController{
		basicAttributes: basicAttributes{
			db:     db,
			logger: logger,
			Gl:     logger,
		},
	}
}

// CreatePost creates a new blog post.
// swagger:route POST /posts CreatePost
// Creates a new blog post.
// responses:
//
//	201: map[string]interface{} "Post created successfully"
//	400: map[string]interface{} "Invalid request"
//	401: map[string]interface{} "Unauthorized"
//	409: map[string]interface{} "Cannot make the post"
//	500: map[string]interface{} "Internal server error"
func (pc *postController) CreatePost(ctx echo.Context) error {
	var post PostRequest
	userID, _ := tools.ExtractUserID(ctx)
	if err := ctx.Bind(&post); err != nil {
		log.Gl.Error(err.Error())
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
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusConflict, echo.Map{"message": "cannot make the post"})
	}

	log.Gl.Info("Post created", zap.Any("post_id", postID))
	return ctx.JSON(http.StatusCreated, echo.Map{
		"message": "Post created successfully",
		"post_id": postID,
	})
}

// DeletePost deletes a blog post.
// swagger:route DELETE /posts/{id} DeletePost
// Deletes a blog post.
// parameters:
//   - name: id
//     in: path
//     description: The ID of the post to delete.
//     required: true
//     type: integer
//
// responses:
//
//	200: map[string]interface{} "Post deleted"
//	401: map[string]interface{} "Unauthorized"
//	500: map[string]interface{} "Internal server error"
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
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	log.Gl.Info("Post deleted", zap.Uint("post_id", uint(postID)))
	return ctx.JSON(http.StatusOK, echo.Map{"message": "Post deleted"})
}

// UpdatePost updates a blog post.
// swagger:route PUT /posts/{id} UpdatePost
// Updates a blog post.
// parameters:
//   - name: id
//     in: path
//     description: The ID of the post to update.
//     required: true
//     type: integer
//
// responses:
//
//	200: map[string]interface{} "Post updated"
//	400: map[string]interface{} "Invalid request"
//	500: map[string]interface{} "Internal server error"
func (pc *postController) UpdatePost(ctx echo.Context) error {
	var post PostRequest
	if err := ctx.Bind(&post); err != nil {
		log.Gl.Error(err.Error())
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
		log.Gl.Error(err.Error())
		return ctx.JSON(500, echo.Map{"error": err.Error()})
	}
	log.Gl.Info("Post updated", zap.Uint("post_id", uint(postID)))
	return ctx.JSON(200, echo.Map{"post": newPost})
}

// GetPost retrieves a blog post by ID.
// swagger:route GET /posts/{id} GetPost
// Retrieves a blog post by ID.
// parameters:
//   - name: id
//     in: path
//     description: The ID of the post to retrieve.
//     required: true
//     type: integer
//
// responses:
//
//	200: map[string]interface{} "Post retrieved"
//	404: map[string]interface{} "Post not found"
//	500: map[string]interface{} "Internal server error"
func (pc *postController) GetPost(ctx echo.Context) error {
	postID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	post, _ := pc.db.GetPost(uint(postID))
	log.Gl.Info("Post retrieved", zap.Uint("post_id", uint(postID)))
	return ctx.JSON(200, echo.Map{"post": post})
}

// GetPosts retrieves all blog posts.
// swagger:route GET /posts GetPosts
// Retrieves all blog posts.
// responses:
//
//	200: map[string]interface{} "Posts retrieved"
//	404: map[string]interface{} "Posts not found"
//	500: map[string]interface{} "Internal server error"
func (pc *postController) GetPosts(ctx echo.Context) error {
	posts, _ := pc.db.GetPosts()
	return ctx.JSON(200, echo.Map{"posts": posts})
}
