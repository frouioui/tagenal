package routes

import "github.com/labstack/echo"

func DefineRouteHandlers(e *echo.Echo) {
	e.GET("/", homeHandler)
	e.GET("/users/", usersHandler)
	e.GET("/users/id/:id", userIDHandler)
	e.GET("/users/region/:region", usersRegionHandler)
}
