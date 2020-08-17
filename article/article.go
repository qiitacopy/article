package article

import "time"

// Article : 記事の構造体
type Article struct {
	ID        int
	Username  string
	Title     string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
