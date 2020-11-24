package routes

import "github.com/labstack/echo"

func DefineRouteHandlers(e *echo.Echo) {
	e.GET("/", homeHandler)
}
