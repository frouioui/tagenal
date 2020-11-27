package server

import (
	"context"
	"log"

	"github.com/frouioui/tagenal/api/articles/db"
	"github.com/frouioui/tagenal/api/articles/pb"
)

type articleServiceGRPC struct {
	pb.UnimplementedArticleServiceServer
	dbm *db.DatabaseManager
}

func newServiceGRPC() (grpcsrv articleServiceGRPC, err error) {
	grpcsrv.dbm, err = db.NewDatabaseManager()
	if err != nil {
		log.Println(err.Error())
		return grpcsrv, err
	}
	return grpcsrv, nil
}

func (s *articleServiceGRPC) GetSingleArticle(cxt context.Context, r *pb.ID) (*pb.Article, error) {
	article, err := s.dbm.GetArticleByID(uint64(r.ID))
	if err != nil {
		return nil, err
	}
	resp := article.ProtoArticle()
	return resp, nil
}

func (s *articleServiceGRPC) GetCategoryArticles(cxt context.Context, r *pb.Category) (*pb.Articles, error) {
	articles, err := s.dbm.GetArticlesOfCategory(r.Category)
	if err != nil {
		return nil, err
	}
	resp := db.ArticlesToProtoArticles(articles)
	return resp, nil
}

func (s *articleServiceGRPC) GetArticlesByRegion(cxt context.Context, r *pb.ID) (*pb.Articles, error) {
	articles, err := s.dbm.GetArticlesFromRegion(int(r.ID))
	if err != nil {
		return nil, err
	}
	resp := db.ArticlesToProtoArticles(articles)
	return resp, nil
}

func (s *articleServiceGRPC) ServiceInformation(cxt context.Context, r *pb.ArticleHomeRequest) (*pb.ArticleHomeResponse, error) {
	resp := &pb.ArticleHomeResponse{}
	resp.IP = getHostIP()
	resp.Host = getHostName()
	return resp, nil
}

func (s *articleServiceGRPC) NewArticle(cxt context.Context, r *pb.Article) (*pb.ID, error) {
	user := db.ProtoArticleToArticle(r)
	id, err := s.dbm.InsertArticle(user)
	if err != nil {
		return nil, err
	}
	pbid := &pb.ID{ID: int64(id)}
	return pbid, nil
}

func (s *articleServiceGRPC) NewArticles(cxt context.Context, r *pb.Articles) (*pb.IDs, error) {
	ids := &pb.IDs{IDs: make([]*pb.ID, 0)}
	for _, u := range r.Articles {
		user := db.ProtoArticleToArticle(u)
		id, err := s.dbm.InsertArticle(user)
		if err != nil {
			return nil, err
		}
		ids.IDs = append(ids.IDs, &pb.ID{ID: int64(id)})
	}
	return ids, nil
}
