package models

type Post struct {
	ID        int     `db:"id"`
	Slug      string  `db:"slug"`
	Title     string  `db:"title"`
	Subtitle  *string `db:"subtitle"` // nullable
	Author    string  `db:"author"`
	Content   string  `db:"content"`
	Tags      *string `db:"tags"`       // nullable
	Category  *string `db:"category"`   // nullable
	CreatedOn string  `db:"created_on"` // treated as string because MariaDB has issues
	UpdatedOn string  `db:"updated_on"` // treated as string because MariaDB has issues
	Published bool    `db:"published"`
}
