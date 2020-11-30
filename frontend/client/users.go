package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/frouioui/tagenal/frontend/models"
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

func UserFromID(ID int) (user *models.User, err error) {
	url := fmt.Sprintf("http://users-api:10000/id/%d", ID)
	method := "GET"

	client := &http.Client{Timeout: time.Second * 10}
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

	var response responseSingleUser
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &response.User, nil
}

func UsersFromRegion(region string) (users []models.User, err error) {
	url := fmt.Sprintf("http://users-api:10000/region/%s", region)
	method := "GET"

	client := &http.Client{Timeout: time.Second * 10}
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

	var response responseArrayUsers
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response.Users, nil
}
