package models

type Author struct {
	Name      string  `db:"id"`
	Location  string  `db:"location"`
	Skills    *string `db:"skills"` // nullable
	Email     string  `db:"email"`
	Website   *string `db:"website"`   // nullable
	Biography *string `db:"biography"` // nullable
	Avatar    *string `db:"avatar"`    // nullable
	Contact   *string `db:"contact"`   // nullable
	Links     *string `db:"links"`     // nullable
	Image     *string `db:"image"`     // nullable

}
