package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/frouioui/tagenal/frontend/client"
	"github.com/labstack/echo/v4"
)

func usersHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}

func usersRegionHandler(c echo.Context) error {
	region := c.Param("region")
	users, err := client.UsersFromRegion(region)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	return c.Render(http.StatusOK, "users_region.htm", map[string]interface{}{
		"page":   "users_region",
		"region": region,
		"users":  users,
	})
}

func userIDHandler(c echo.Context) error {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusOK, err.Error())
	}
	user, err := client.UserFromID(ID)
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusOK, err.Error())
	}
	log.Println(user)
	return c.Render(http.StatusOK, "user.htm", map[string]interface{}{
		"page": "user_id",
		"id":   ID,
		"user": user,
	})
}
