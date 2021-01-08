package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	pb "github.com/frouioui/tagenal/frontend/client/pb/articles"

	"github.com/frouioui/tagenal/frontend/models"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	otlog "github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type responseSingleArticle struct {
	Status  string         `json:"status"`
	Code    int            `json:"code"`
	Article models.Article `json:"data"`
}

type responseArrayArticles struct {
	Status   string           `json:"status"`
	Code     int              `json:"code"`
	Articles []models.Article `json:"data"`
}

var grpcArticlesClient pb.ArticleServiceClient

func InitArticlesGRPC() (err error) {
	tracer := opentracing.GlobalTracer()
	conn, err := grpc.Dial(os.Getenv("ARTICLES_API_SERVICE_HOST")+":"+os.Getenv("ARTICLES_API_SERVICE_PORT_GRPC"),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)),
		grpc.WithStreamInterceptor(otgrpc.OpenTracingStreamClientInterceptor(tracer)),
	)
	grpcArticlesClient = pb.NewArticleServiceClient(conn)
	return err
}

func ArticlesFromIDGRPC(c echo.Context, ID int) (article *models.Article, err error) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second)
	defer cancel()
	r, err := grpcArticlesClient.GetSingleArticle(ctx, &pb.ID{ID: int64(ID)})
	if err != nil {
		s := status.Convert(err)
		switch s.Code() {
		case codes.NotFound:
			log.Printf("User not found: %s", s.Message())
		default:
			log.Printf("Error: %s", s.Message())
		}
		return nil, err
	}
	return &models.Article{
		ID:          r.ID,
		Timestamp:   r.Timestamp,
		AID:         r.AID,
		Title:       r.Title,
		Category:    r.Category,
		Abstract:    r.Abstract,
		ArticleTags: r.ArticleTags,
		Authors:     r.Authors,
		Language:    r.Language,
		Text:        r.Text,
		Image:       r.Image,
		Video:       r.Video,
	}, nil
}

func ArticleFromID(c echo.Context, ID int) (article *models.Article, err error) {
	url := fmt.Sprintf("http://articles-api:8080/id/%d", ID)
	method := "GET"

	client := &http.Client{Timeout: time.Second * 10}

	span := jaegertracing.CreateChildSpan(c, "ArticleFromID")
	defer span.Finish()
	req, err := jaegertracing.NewTracedRequest(method, url, nil, span)

	if err != nil {
		log.Println(err.Error())
		span.LogFields(otlog.String("err", err.Error()))
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		span.LogFields(otlog.String("err", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	var response responseSingleArticle
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println(err.Error())
		span.LogFields(otlog.String("err", err.Error()))
		return nil, err
	}
	span.LogFields(
		otlog.String("event", "api call ArticleFromID"),
		otlog.String("value", response.Status),
	)
	return &response.Article, nil
}

func ArticlesFromCategoryGRPC(c echo.Context, category string) (articles []models.Article, err error) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second)
	defer cancel()
	r, err := grpcArticlesClient.GetCategoryArticles(ctx, &pb.Category{Category: category})
	if err != nil {
		s := status.Convert(err)
		switch s.Code() {
		case codes.NotFound:
			log.Printf("User not found: %s", s.Message())
		default:
			log.Printf("Error: %s", s.Message())
		}
		return nil, err
	}
	return models.TransformProtosToArticles(r), nil
}

func ArticlesFromCategory(c echo.Context, category string) (articles []models.Article, err error) {
	url := fmt.Sprintf("http://articles-api:8080/category/%s", category)
	method := "GET"

	client := &http.Client{Timeout: time.Second * 10}

	span := jaegertracing.CreateChildSpan(c, "ArticlesFromCategory")
	defer span.Finish()
	req, err := jaegertracing.NewTracedRequest(method, url, nil, span)

	if err != nil {
		log.Println(err.Error())
		span.LogFields(otlog.String("err", err.Error()))
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		span.LogFields(otlog.String("err", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	var response responseArrayArticles
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println(err.Error())
		span.LogFields(otlog.String("err", err.Error()))
		return nil, err
	}
	span.LogFields(
		otlog.String("event", "api call ArticlesFromCategory"),
		otlog.String("value", response.Status),
	)
	return response.Articles, nil
}

func ArticlesFromRegionGRPC(c echo.Context, regionID int) (articles []models.Article, err error) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second)
	defer cancel()
	r, err := grpcArticlesClient.GetArticlesByRegion(ctx, &pb.ID{ID: int64(regionID)})
	if err != nil {
		s := status.Convert(err)
		switch s.Code() {
		case codes.NotFound:
			log.Printf("User not found: %s", s.Message())
		default:
			log.Printf("Error: %s", s.Message())
		}
		return nil, err
	}
	return models.TransformProtosToArticles(r), nil
}

func ArticlesFromRegion(c echo.Context, regionID int) (articles []models.Article, err error) {
	url := fmt.Sprintf("http://articles-api:8080/region/id/%d", regionID)
	method := "GET"

	client := &http.Client{Timeout: time.Second * 10}

	span := jaegertracing.CreateChildSpan(c, "ArticleFromRegion")
	defer span.Finish()
	req, err := jaegertracing.NewTracedRequest(method, url, nil, span)

	if err != nil {
		log.Println(err.Error())
		span.LogFields(otlog.String("err", err.Error()))
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		span.LogFields(otlog.String("err", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	var response responseArrayArticles
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println(err.Error())
		span.LogFields(otlog.String("err", err.Error()))
		return nil, err
	}
	span.LogFields(
		otlog.String("event", "api call ArticleFromRegion"),
		otlog.String("value", response.Status),
	)
	return response.Articles, nil
}
