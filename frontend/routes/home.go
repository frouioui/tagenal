package routes

import (
	"net/http"

	"github.com/frouioui/tagenal/frontend/models"
	"github.com/labstack/echo/v4"
)

func homeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.htm", map[string]interface{}{
		"page":     "home",
		"srvinfos": models.GetDefaultServicesInfos(),
	})
}
