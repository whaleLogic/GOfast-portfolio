package handlers

import (
	"blog/models"
	"database/sql"
	"fmt"
)

func GetPostBySlug(slug string) (models.Post, error) {
	dsn := "root:Sophia@tcp(127.0.0.1:3306)/content"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return models.Post{}, fmt.Errorf("DB open failed: %w", err)
	}
	defer db.Close()

	var p models.Post
	err = db.QueryRow("SELECT * FROM all_posts WHERE slug = ?", slug).Scan(
		&p.ID,
		&p.Slug,
		&p.Title,
		&p.Subtitle,
		&p.Author,
		&p.Content,
		&p.Tags,
		&p.Category,
		&p.CreatedOn,
		&p.UpdatedOn,
		&p.Published,
	)
	if err != nil {
		return models.Post{}, fmt.Errorf("Query failed: %w", err)
	}
	return p, nil
}

