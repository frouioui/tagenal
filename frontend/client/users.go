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

type responseSingleUser struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	User   models.User `json:"data"`
}

type responseArrayUsers struct {
	Status string        `json:"status"`
	Code   int           `json:"code"`
	Users  []models.User `json:"data"`
}

func UserFromID(c echo.Context, ID int) (user *models.User, err error) {
	url := fmt.Sprintf("http://users-api:10000/id/%d", ID)
	method := "GET"

	client := &http.Client{Timeout: time.Second * 10}

	span := jaegertracing.CreateChildSpan(c, "UserFromID")
	defer span.Finish()
	req, err := jaegertracing.NewTracedRequest(method, url, nil, span)

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

	var response responseSingleUser
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	span.LogFields(
		otlog.String("event", "api call UserFromID"),
		otlog.String("value", response.Status),
	)
	return &response.User, nil
}

func UsersFromRegion(c echo.Context, region string) (users []models.User, err error) {
	url := fmt.Sprintf("http://users-api:10000/region/%s", region)
	method := "GET"

	client := &http.Client{Timeout: time.Second * 10}
	span := jaegertracing.CreateChildSpan(c, "UsersFromRegion")
	defer span.Finish()
	req, err := jaegertracing.NewTracedRequest(method, url, nil, span)

	if err != nil {
		span.LogFields(otlog.String("err", err.Error()))
		log.Println(err.Error())
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		span.LogFields(otlog.String("err", err.Error()))
		log.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	var response responseArrayUsers
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		span.LogFields(otlog.String("err", err.Error()))
		log.Println(err.Error())
		return nil, err
	}
	span.LogFields(
		otlog.String("event", "api call UsersFromRegion"),
		otlog.String("value", response.Status),
	)
	return response.Users, nil
}
