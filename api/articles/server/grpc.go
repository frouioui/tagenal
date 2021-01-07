package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func (s *articleServiceGRPC) ServiceInformation(ctx context.Context, r *pb.InformationRequest) (*pb.InformationResponse, error) {
	resp := &pb.InformationResponse{}
	resp.IP = getHostIP()
	resp.Host = getHostName()
	return resp, nil
}

func (s *articleServiceGRPC) GetSingleArticle(ctx context.Context, r *pb.ID) (*pb.Article, error) {
	var article db.Article
	var err error
	article, err = getCacheArticle(ctx, fmt.Sprintf("article_id_%d", r.ID), article)
	if err != nil {
		article, err := s.dbm.GetArticleByID(ctx, "", uint64(r.ID))
		if err != nil {
			if err == sql.ErrNoRows {
				st := status.New(codes.NotFound, "Not found.")
				return nil, st.Err()
			}
			st, _ := status.FromError(err)
			return nil, st.Err()
		}
		err = setCacheArticle(ctx, fmt.Sprintf("article_id_%d", r.ID), article)
		if err != nil {
			log.Println(err.Error())
			st, _ := status.FromError(err)
			return nil, st.Err()
		}
	}
	return article.ProtoArticle(), nil
}

func (s *articleServiceGRPC) GetCategoryArticles(ctx context.Context, r *pb.Category) (*pb.Articles, error) {
	articles, err := s.dbm.GetArticlesOfCategory(ctx, "", r.Category)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	resp := db.ArticlesToProtoArticles(articles)
	return resp, nil
}

func (s *articleServiceGRPC) GetArticlesByRegion(ctx context.Context, r *pb.ID) (*pb.Articles, error) {
	articles, err := s.dbm.GetArticlesFromRegion(ctx, "", int(r.ID))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	resp := db.ArticlesToProtoArticles(articles)
	return resp, nil
}

func (s *articleServiceGRPC) NewArticle(ctx context.Context, r *pb.Article) (*pb.ID, error) {
	user := db.ProtoArticleToArticle(r)
	id, err := s.dbm.InsertArticle(user)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	pbid := &pb.ID{ID: int64(id)}
	return pbid, nil
}

func (s *articleServiceGRPC) NewArticles(ctx context.Context, r *pb.Articles) (*pb.IDs, error) {
	ids := &pb.IDs{IDs: make([]*pb.ID, 0)}
	for _, u := range r.Articles {
		user := db.ProtoArticleToArticle(u)
		id, err := s.dbm.InsertArticle(user)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		ids.IDs = append(ids.IDs, &pb.ID{ID: int64(id)})
	}
	return ids, nil
}
