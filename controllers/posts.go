package controllers

import (
	"strconv"

	"github.com/arshamalh/blogo/database"
	"github.com/arshamalh/blogo/models"
	"github.com/gin-gonic/gin"
)

type PostRequest struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
	var post PostRequest
	value, _ := c.Get("user_id")
	if c.BindJSON(&post) == nil {
		new_post := models.Post{
			Title:    post.Title,
			Content:  post.Content,
			AuthorID: uint(value.(int)),
		}
		post_id, _ := database.CreatePost(&new_post)
		c.JSON(200, gin.H{
			"message": "Post created successfully",
			"post_id": post_id,
		})
	}
}

func GetPost(c *gin.Context) {
	post_id, _ := strconv.Atoi(c.Param("id"))
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
