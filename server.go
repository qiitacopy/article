package main

import (
	"log"
	"net"

	pb "github.com/qiitacopy/article/grpc"
	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

func main() {

	// リッスン処理
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// サーバ起動
	s := grpc.NewServer()
	pb.RegisterArticleServiceServer(s, &pb.ArticleGrpcServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
