package controllers

import (
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
	user_id, err := tools.ExtractUserID(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
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

func GetPost(c *gin.Context) {
	post_id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	post, _ := database.GetPost(uint(post_id))
	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	posts, _ := database.GetPosts()
	c.JSON(200, gin.H{
		"posts": posts,
	})
}
