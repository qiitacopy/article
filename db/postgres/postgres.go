package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/qiitacopy/article/article"

	// GORMが公開するPostgres公式用ドライバー
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Postgres : Postgresのデータベース接続
type Postgres struct{}

// NewPostgres : Postgresのコンストラクタ
func NewPostgres() *Postgres {
	return new(Postgres)
}

// GetByID : 主キー検索
func (p *Postgres) GetByID(id int) (*article.Article, error) {
	// DB接続
	db, err := gormConnect()
	// エラー処理
	if err != nil {
		return nil, err
	}
	defer db.Close()
	// 一件取得
	article := new(article.Article)
	article.ID = id
	db.First(&article)

	return article, nil
}

// CreateArticle : 一件登録
func (p *Postgres) CreateArticle(article *article.Article) (*article.Article, error) {
	// DB接続
	db, err := gormConnect()
	// エラー処理
	if err != nil {
		return nil, err
	}
	defer db.Close()
	// １件登録 CreatedAt,UpdatedAtはGORMの機能で自動で入力される
	db.Create(&article)
	// 登録内容の取得
	db.First(&article)
	return article, nil
}

// gromConnect : postgresに接続する
func gormConnect() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "postgres://test:test@postgres:5432/qiita?sslmode=disable")
	return db, err
}
