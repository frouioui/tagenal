package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func homeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}
