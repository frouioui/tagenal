package routes

import "github.com/labstack/echo/v4"

func DefineRouteHandlers(e *echo.Echo) {
	e.GET("/", homeHandler)
	e.GET("/users/", usersHandler)
	e.GET("/users/id/:id", userIDHandler)
	e.GET("/users/region/:region", usersRegionHandler)
	e.GET("/articles/", articlesHandler)
	e.GET("/articles/id/:id", articleIDHandler)
	e.GET("/articles/category/:category", articlesCategoryHandler)
	e.GET("/articles/region/:region", articlesRegionHandler)
	e.GET("/status", servicesInfoHandler)
}
