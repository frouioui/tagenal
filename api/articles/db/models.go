package db

import (
	"encoding/json"

	"github.com/frouioui/tagenal/api/articles/pb"
)

// ArticleArray represents an array of Article.
// Implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler.
type ArticleArray []Article

// Article struct refers to the article table of the
// Vitess MySQl cluster, in the articles keyspace.
type Article struct {
	ID          int64  `json:"id"`
	Timestamp   int64  `json:"timestamp"`
	AID         string `json:"aid"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Abstract    string `json:"abstract"`
	ArticleTags string `json:"article_tags"`
	Authors     string `json:"authors"`
	Language    string `json:"language"`
	Text        string `json:"text"`
	Image       string `json:"image"`
	Video       string `json:"video"`
}

// ProtoArticle transforms an Article into a the auto-generated
// pb.Article structure from protobuf.
func (u *Article) ProtoArticle() *pb.Article {
	return &pb.Article{
		ID:          u.ID,
		Timestamp:   u.Timestamp,
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

// ProtoArticleToArticle transforms an auto-generated pb.Article from
// protobuf into the package implementation of Article.
func ProtoArticleToArticle(pbarticle *pb.Article) (article Article) {
	return Article{
		ID:          pbarticle.ID,
		Timestamp:   pbarticle.Timestamp,
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

// ArticlesToProtoArticles transforms an array of Articles into
// an array of pb.Articles which are auto-generated from protobuf.
func ArticlesToProtoArticles(articles []Article) *pb.Articles {
	pbarticles := &pb.Articles{Articles: make([]*pb.Article, len(articles))}
	for i, a := range articles {
		pbarticles.Articles[i] = a.ProtoArticle()
	}
	return pbarticles
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (u *Article) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &u)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (u *Article) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (ua *ArticleArray) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &ua)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (ua *ArticleArray) MarshalBinary() (data []byte, err error) {
	return json.Marshal(ua)
}
