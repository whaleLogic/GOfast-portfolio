package handlers

import (
	"blog/models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllPosts() ([]models.Post, error) {
	dsn := "root:Sophia@tcp(127.0.0.1:3306)/content"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}
	defer db.Close()

	fmt.Printf("Connected to DB: %T\n", db)

	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, fmt.Errorf("failed to query DB: %w", err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(
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
		); err != nil {
			return nil, fmt.Errorf("row scan error: %w", err)
		}
		posts = append(posts, p)
	}
	return posts, nil
}
