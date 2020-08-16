package model

import (
	"time"
)

// Article : 記事の構造体
type Article struct {
	ID         int
	Title      string
	Content    string
	Author     string
	PostDate   time.Time
	UpdateDate time.Time
}
