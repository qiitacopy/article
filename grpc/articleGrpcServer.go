package grpc

import (
	context "context"
	"log"
)

// ArticleGrpcServer : 自動生成されたArticleServiceServerインターフェースの実装
type ArticleGrpcServer struct{}

// GetByID : 記事IDに該当する記事を１件取得する
func (s ArticleGrpcServer) GetByID(ctx context.Context, articleID *ArticleID) (*Article, error) {
	log.Printf("Received: %v", articleID.Id)
	// TODO DBからarticle.Articleを一件取得
	// TODO article.Articleからgrpc.Articleに変換
	article := new(Article)
	article.Id = articleID.Id
	article.Title = "タイトル"
	article.Text = "本文"
	article.Username = "著者"
	return article, nil
}
