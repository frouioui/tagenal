package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func homeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.htm", "")
}
