package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/frouioui/tagenal/frontend/models"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	otlog "github.com/opentracing/opentracing-go/log"
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

func ArticleFromCategory(c echo.Context, category string) (articles []models.Article, err error) {
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

func ArticleFromRegion(c echo.Context, regionID int) (articles []models.Article, err error) {
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
