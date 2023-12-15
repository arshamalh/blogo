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
		},
	}
}

// @Summary Create a new blog post
// @Description Create a new blog post
// @ID create-post
// @Accept json
// @Produce json
// @Param post body PostRequest true "Post object"
// @Success 201 {object} map[string]interface{} "Post created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 409 {object} map[string]interface{} "Cannot make the post"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /posts [post]
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

// @Summary Delete a blog post
// @Description Delete a blog post
// @ID delete-post
// @Param id path int true "Post ID"
// @Success 200 {object} map[string]interface{} "Post deleted"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /posts/{id} [delete]
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

// @Summary Update a blog post
// @Description Update a blog post
// @ID update-post
// @Param id path int true "Post ID"
// @Accept json
// @Produce json
// @Param post body PostRequest true "Post object"
// @Success 200 {object} map[string]interface{} "Post updated"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /posts/{id} [put]
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

// @Summary Get a blog post by ID
// @Description Get a blog post by ID
// @ID get-post
// @Param id path int true "Post ID"
// @Produce json
// @Success 200 {object} map[string]interface{} "Post retrieved"
// @Failure 404 {object} map[string]interface{} "Post not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /posts/{id} [get]
func (pc *postController) GetPost(ctx echo.Context) error {
	postID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	post, _ := pc.db.GetPost(uint(postID))
	log.Gl.Info("Post retrieved", zap.Uint("post_id", uint(postID)))
	return ctx.JSON(200, echo.Map{"post": post})
}

// @Summary Get all blog posts
// @Description Get all blog posts
// @ID get-posts
// @Produce json
// @Success 200 {object} map[string]interface{} "Posts retrieved"
// @Failure 404 {object} map[string]interface{} "Posts not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /posts [get]
func (pc *postController) GetPosts(ctx echo.Context) error {
	posts, _ := pc.db.GetPosts()
	return ctx.JSON(200, echo.Map{"posts": posts})
}
