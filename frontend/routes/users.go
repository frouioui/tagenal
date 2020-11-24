package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func usersHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}

func usersRegionHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}

func userIDHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}
