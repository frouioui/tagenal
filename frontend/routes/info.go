package routes

import (
	"net/http"

	"github.com/frouioui/tagenal/frontend/models"

	"github.com/labstack/echo"
)

func servicesInfoHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "info.htm", map[string]interface{}{
		"page":     "status",
		"srvinfos": models.GetDefaultServicesInfos(),
	})
}
