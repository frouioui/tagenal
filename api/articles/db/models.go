package db

import (
	"github.com/frouioui/tagenal/api/articles/pb"
)

type Article struct {
	ID          int64  `json:"ID"`
	Timestamp   string `json:"Timestamp"`
	ID2         string `json:"ID2"`
	AID         string `json:"AID"`
	Title       string `json:"Title"`
	Category    string `json:"Category"`
	Abstract    string `json:"Abstract"`
	ArticleTags string `json:"ArticleTags"`
	Authors     string `json:"Authors"`
	Language    string `json:"Language"`
	Text        string `json:"Text"`
	Image       string `json:"Image"`
	Video       string `json:"Video"`
}

func (u *Article) ProtoArticle() *pb.Article {
	return &pb.Article{
		ID:          u.ID,
		Timestamp:   u.Timestamp,
		ID2:         u.ID2,
		AID:         u.AID,
		Title:       u.Title,
		Category:    u.Category,
		Abstract:    u.Abstract,
		ArticleTags: u.ArticleTags,
		Authors:     u.Authors,
		Language:    u.Language,
		Text:        u.Text,
		Image:       u.Image,
		Video:       u.Video,
	}
}

func ProtoArticleToArticle(pbarticle *pb.Article) (article Article) {
	return Article{
		ID:          pbarticle.ID,
		Timestamp:   pbarticle.Timestamp,
		ID2:         pbarticle.ID2,
		AID:         pbarticle.AID,
		Title:       pbarticle.Title,
		Category:    pbarticle.Category,
		Abstract:    pbarticle.Abstract,
		ArticleTags: pbarticle.ArticleTags,
		Authors:     pbarticle.Authors,
		Language:    pbarticle.Language,
		Text:        pbarticle.Text,
		Image:       pbarticle.Image,
		Video:       pbarticle.Video,
	}
}

func ArticlesToProtoArticles(articles []Article) *pb.Articles {
	pbarticles := &pb.Articles{Articles: make([]*pb.Article, len(articles))}
	for i, a := range articles {
		pbarticles.Articles[i] = a.ProtoArticle()
	}
	return pbarticles
}
