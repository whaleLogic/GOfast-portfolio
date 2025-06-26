package main

import (
	"blog/handlers"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// database
	posts, err := handlers.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Printf("Exctracted %di Posts", len(posts))

	// 1️⃣ Define your template functions first...
	funcs := template.FuncMap{
		"eq": func(a, b string) bool { return a == b },
		"safeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
	}

	// Parse base first, then every page
	tmpl := template.Must(
		template.New("").Funcs(funcs).ParseFiles(
			"templates/index.html",
			"templates/about.html",
			"templates/projects.html",
			"templates/post.html",
			"templates/blogposts.html",
			"templates/header.html",
			"templates/footer.html",
		),
	)
	// 3️⃣ Tell Gin to use *only* that parsed template set
	r.SetHTMLTemplate(tmpl)

	// 4️⃣ Your routes
	r.GET("/", handlers.HomeHandler)
	r.GET("/about", handlers.AboutHandler)
	r.GET("/projects", handlers.ProjectsHandler)
	r.GET("/blogposts", handlers.BlogPostsHandler)
	r.GET("/post/:slug", handlers.PostHandler)
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}
