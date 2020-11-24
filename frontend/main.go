package main

import (
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/frouioui/tagenal/frontend/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func getTemplateEngine() *Template {
	tmpl := &Template{}
	tmpl.templates = template.New("")
	err := filepath.Walk("./templates", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".htm") {
			_, err = tmpl.templates.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}
		return err
	})
	if err != nil {
		panic(err)
	}
	return tmpl
}

func main() {
	e := echo.New()

	e.Renderer = getTemplateEngine()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "assets")

	routes.DefineRouteHandlers(e)

	e.Logger.Fatal(e.Start(":80"))
}
