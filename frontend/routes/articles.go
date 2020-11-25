package routes

import (
	"log"
	"net/http"

	"github.com/frouioui/tagenal/frontend/client"
	"github.com/labstack/echo"
)

func articlesHandler(c echo.Context) error {
	ar, err := client.ArticleFromCategory("science")
	if err != nil {
		log.Println(err)
	}
	log.Println("articles:", ar)
	return c.String(http.StatusOK, "hello")
}

func articlesCategoryHandler(c echo.Context) error {
	category := c.Param("category")
	ars, err := client.ArticleFromCategory(category)
	if err != nil {
		return c.String(http.StatusOK, "hello")
	}
	return c.Render(http.StatusOK, "articles_category.htm", map[string]interface{}{
		"category": category,
		"articles": ars,
	})
}

func articleIDHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}
