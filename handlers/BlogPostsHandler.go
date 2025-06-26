package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BlogPostsHandler(c *gin.Context) {
	posts, _ := GetAllPosts() // your DB call
	c.HTML(http.StatusOK, "blogposts", gin.H{
		"Page":  "Blog",
		"Posts": posts,
	})
}
