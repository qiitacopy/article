package db

import "github.com/qiitacopy/article/article"

// Database : データベースインターフェース
type Database interface {
	GetByID(int) (*article.Article, error)
}
