package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AboutHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{
		"Title":    "About",
		"Author":   "Keith Thomson",
		"Location": "Connecticut, USA",
		"Skills":   "Python, Golang, SQL, Cloud, Docker",
	})
}
