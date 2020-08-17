package postgres

import "github.com/qiitacopy/article/article"

// Postgres : Postgresのデータベース接続
type Postgres struct{}

// GetById : 主キー検索
func (p *Postgres) GetByID(int) (*article.Article, error) {
	return nil, nil
}
