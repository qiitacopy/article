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
func (s ArticleGrpcServer) GetByID(ctx context.Context, articleID *GetByIDRequest) (*Article, error) {
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

// CreateArticle : 記事を１件登録する
func (s ArticleGrpcServer) CreateArticle(ctx context.Context, request *CreateArticleRequest) (*Article, error) {
	log.Printf("gRPC: CreateArticle is called")
	// CreateArticleRequestからarticle.Articleに変換
	article := convertCreateArticleRequestToRecord(request)
	// １件登録
	record, err := s.db.CreateArticle(article)
	// DB接続でエラーが発生した場合、呼び出しもとに戻す
	if err != nil {
		return nil, err
	}
	// DB用記事モデルをgRPC用記事モデルに変換
	response := convertArticleRecordToGrpc(record)
	return response, nil
}

// convertCreateArticleRequestToRecord : DB用記事モデルに変換
func convertCreateArticleRequestToRecord(request *CreateArticleRequest) *article.Article {
	article := new(article.Article)
	article.Title = request.Title
	article.Text = request.Text
	article.Username = request.Username
	return article
}

// convertArticleRecordToGrpc : DB用記事モデルをgRPC用記事モデルに変換
func convertArticleRecordToGrpc(record *article.Article) *Article {
	const timeFormat = "2006/01/02 15:04:05"
	article := new(Article)
	article.Id = int32(record.ID)
	article.Title = record.Title
	article.Text = record.Text
	article.Username = record.Username
	article.CreatedAt = record.CreatedAt.Format(timeFormat)
	article.UpdatedAt = record.UpdatedAt.Format(timeFormat)
	return article
}
