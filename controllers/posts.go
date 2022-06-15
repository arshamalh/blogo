package controllers

import (
	"net/http"
	"strconv"

	"github.com/arshamalh/blogo/database"
	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/tools"
	"github.com/gin-gonic/gin"
)

type PostRequest struct {
	Title      string   `form:"title" json:"title" binding:"required"`
	Content    string   `form:"content" json:"content" binding:"required"`
	Categories []string `form:"categories" json:"categories"`
}

func CreatePost(ctx *gin.Context) {
	var post PostRequest
	user_id, _ := tools.ExtractUserID(ctx)
	if err := ctx.BindJSON(&post); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
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
	post_id, _ := database.CreatePost(&new_post)
	ctx.JSON(200, gin.H{
		"message": "Post created successfully",
		"post_id": post_id,
	})
}

func DeletePost(ctx *gin.Context) {
	user_id, _ := tools.ExtractUserID(ctx)
	post_id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	// If user doesn't have the permission to delete posts, check if it's its own post.
	if !tools.ExtractPermissable(ctx) {
		post, _ := database.GetPost(uint(post_id))
		if post.AuthorID != user_id {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "you don't have enough permissions",
			})
			return
		}
	}

	// Delete post
	if err := database.DeletePost(uint(post_id)); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Post deleted"})
}

func GetPost(ctx *gin.Context) {
	post_id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	post, _ := database.GetPost(uint(post_id))
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(ctx *gin.Context) {
	posts, _ := database.GetPosts()
	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}
