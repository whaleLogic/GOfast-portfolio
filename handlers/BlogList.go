package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"html/template"
	"net/http"
)

func BlogPostHandler(c *gin.Context) {
	slug := c.Param("slug")

	post, err := GetPostBySlug(slug)
	if err != nil {
		c.String(http.StatusNotFound, "Post not found")
		return
	}

	rendered := markdown.ToHTML([]byte(post.Content), nil, nil)

	c.HTML(http.StatusOK, "base.html", gin.H{
		"Title": post.Title,
		"HTML":  template.HTML(rendered),
		"Author": post.Author,
		"CreatedOn": post.CreatedOn,
	})
}

