package main

import (
	"github.com/frouioui/tagenal/frontend/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.DefineRouteHandlers(e)

	e.Logger.Fatal(e.Start(":80"))
}
