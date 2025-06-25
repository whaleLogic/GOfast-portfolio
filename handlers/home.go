package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "Home",
	})
}

func AboutHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{
		"Title": "About",
	})
}

func ProjectsHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "projects.html", gin.H{
		"Title": "Projects",
	})
}

