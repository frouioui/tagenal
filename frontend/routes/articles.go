package routes

import (
	"log"
	"net/http"
	"strconv"

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
		return c.String(http.StatusOK, err.Error())
	}
	return c.Render(http.StatusOK, "articles_category.htm", map[string]interface{}{
		"page":     "articles_category",
		"category": category,
		"articles": ars,
	})
}

func articlesRegionHandler(c echo.Context) error {
	region := c.Param("region")
	regionID := 1
	// TODO: create a global map for region id
	if region == "beijing" {
		regionID = 1
	} else if region == "hong kong" {
		regionID = 2
	}
	ars, err := client.ArticleFromRegion(regionID)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	return c.Render(http.StatusOK, "articles_region.htm", map[string]interface{}{
		"page":     "articles_region",
		"region":   region,
		"articles": ars,
	})
}

func articleIDHandler(c echo.Context) error {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusOK, err.Error())
	}
	art, err := client.ArticleFromID(ID)
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusOK, err.Error())
	}
	return c.Render(http.StatusOK, "article.htm", map[string]interface{}{
		"page":    "article_id",
		"id":      ID,
		"article": art,
	})
}
