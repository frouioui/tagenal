package routes

import (
	"net/http"

	"github.com/frouioui/tagenal/frontend/models"

	"github.com/labstack/echo/v4"
)

func servicesInfoHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "info.htm", map[string]interface{}{
		"page":     "status",
		"srvinfos": models.GetDefaultServicesInfos(),
	})
}

func healthHandler(c echo.Context) error {
	return c.JSON(200, `ok`)
}

func readyHandler(c echo.Context) error {
	return c.JSON(200, `ok`)
}
