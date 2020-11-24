package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/frouioui/tagenal/frontend/models"
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

func ArticleFromID(ID int) (article *models.Article, err error) {
	url := fmt.Sprintf("http://articles-api-service:8080/id/%d", ID)
	method := "GET"

	client := &http.Client{Timeout: time.Second * 2}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	var response responseSingleArticle
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &response.Article, nil
}

func ArticleFromCategory(category string) (articles []models.Article, err error) {
	url := fmt.Sprintf("http://articles-api:8080/category/%s", category)
	method := "GET"

	client := &http.Client{Timeout: time.Second * 2}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	var response responseArrayArticles
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response.Articles, nil
}
