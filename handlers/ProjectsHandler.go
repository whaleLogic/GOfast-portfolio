package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProjectsHandler(c *gin.Context) {
	p, err := GetAllPosts()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching projects: %v", err)
		return
	}
	c.HTML(http.StatusOK, "projects.html", gin.H{
		"Title":       "Projects",
		"ProjectList": p,
	})
}
