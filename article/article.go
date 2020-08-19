package article

import "time"

// Article : 記事の構造体
type Article struct {
	ID        int       `gorm:"column:id;primary_key;atuo_increment"`
	Username  string    `gorm:"column:username"`
	Title     string    `gorm:"column:title"`
	Text      string    `gorm:"column:text"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
