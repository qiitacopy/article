package repository

import (
	"github.com/qiitacopy/article/domain/model"
)

// ArticleRepository : Articleのリポジトリインターフェース
type ArticleRepository interface {
	FindById(int) (*model.Article, error)
}
