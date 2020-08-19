package grpc

import (
	context "context"
	"log"

	"github.com/qiitacopy/article/article"
	"github.com/qiitacopy/article/db"
)

// ArticleGrpcServer : 自動生成されたArticleServiceServerインターフェースの実装
type ArticleGrpcServer struct {
	db db.Database
}

// NewArticleServer : ArticleServerのコンストラクタ
func NewArticleServer(db db.Database) *ArticleGrpcServer {
	server := new(ArticleGrpcServer)
	server.db = db
	return server
}

// GetByID : 記事IDに該当する記事を１件取得する
func (s ArticleGrpcServer) GetByID(ctx context.Context, articleID *ArticleID) (*Article, error) {
	log.Printf("gRPC: GetByID is called, ArticleID = %v", articleID.Id)
	// DBからarticle.Articleを一件取得
	record, err := s.db.GetByID(int(articleID.Id))
	// DB接続でエラーが発生した場合、呼び出しもとに戻す
	if err != nil {
		return nil, err
	}
	// article.Articleからgrpc.Articleに変換
	article := convertArticleRecordToGrpc(record)
	return article, nil
}

// convertArticleRecordToGrpc : DB用記事モデルをgRPC用記事モデルに変換
func convertArticleRecordToGrpc(record *article.Article) *Article {
	article := new(Article)
	article.Id = int32(record.ID)
	article.Title = record.Title
	article.Text = record.Text
	article.Username = record.Username
	article.CreatedAt = record.CreatedAt.Format("2019/03/01 10:00:00")
	article.UpdatedAt = record.UpdatedAt.Format("2019/03/01 10:00:00")
	return article
}
