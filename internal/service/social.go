package service

import (
	"context"

	pb "kratos-test/api/realworld/v1"
)


func (s *RealworldService) GetArticle(ctx context.Context, req *pb.GetArticleReq) (*pb.ArticleReply, error) {
	return &pb.ArticleReply{}, nil
}
func (s *RealworldService) GetArticles(ctx context.Context, req *pb.ListArticles) (*pb.MultipleArticlesReply, error) {
	return &pb.MultipleArticlesReply{}, nil
}
func (s *RealworldService) FeedArticles(ctx context.Context, req *pb.FeedArticlesRequest) (*pb.MultipleArticlesReply, error) {
	return &pb.MultipleArticlesReply{}, nil
}
func (s *RealworldService) CreateArticle(ctx context.Context, req *pb.CreateArticleReq) (*pb.ArticleReply, error) {
	return &pb.ArticleReply{}, nil
}
func (s *RealworldService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleReq) (*pb.ArticleReply, error) {
	return &pb.ArticleReply{}, nil
}
func (s *RealworldService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleReq) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *RealworldService) AddComment(ctx context.Context, req *pb.AddCommentReq) (*pb.CommentReply, error) {
	return &pb.CommentReply{}, nil
}
func (s *RealworldService) GetComments(ctx context.Context, req *pb.GetCommentsReq) (*pb.MultipleArticlesReply, error) {
	return &pb.MultipleArticlesReply{}, nil
}
func (s *RealworldService) DeleteComment(ctx context.Context, req *pb.DeleteCommentReq) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *RealworldService) FavoriteArticle(ctx context.Context, req *pb.FavoriteArticleReq) (*pb.ArticleReply, error) {
	return &pb.ArticleReply{}, nil
}
func (s *RealworldService) UnfavoriteArticle(ctx context.Context, req *pb.FavoriteArticleReq) (*pb.ArticleReply, error) {
	return &pb.ArticleReply{}, nil
}

func (s *RealworldService) GetTags(ctx context.Context, req *pb.Empty) (*pb.TagsReply, error) {
	return &pb.TagsReply{}, nil
}
