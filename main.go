package main

import (
	"blog/handlers"
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Template functions
	r.SetFuncMap(template.FuncMap{
		"eq": func(a, b string) bool { return a == b },
		"safeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
	})

	// Load templates after defining FuncMap
	r.LoadHTMLGlob("templates/*.html")

	// Route setup
	r.GET("/", handlers.HomeHandler)
	r.GET("/about", handlers.AboutHandler)
	r.GET("/projects", handlers.ProjectsHandler)
	r.GET("/post/:slug", handlers.BlogPostHandler)

	log.Println("Starting server on http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}

