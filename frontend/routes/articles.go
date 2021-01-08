package routes

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/frouioui/tagenal/frontend/client"
	"github.com/labstack/echo/v4"
)

func articlesHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}

func articlesCategoryHandler(c echo.Context) error {
	category := c.Param("category")
	ars, err := client.ArticlesFromCategoryGRPC(c, category)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	for i, ar := range ars {
		imgs, _ := ar.GetAssetsInfo()
		if len(imgs) > 0 {
			ars[i].Image = imgs[0]
		}
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
	ars, err := client.ArticlesFromRegionGRPC(c, regionID)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	for i, ar := range ars {
		imgs, _ := ar.GetAssetsInfo()
		if len(imgs) > 0 {
			ars[i].Image = imgs[0]
		}
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
	art, err := client.ArticlesFromIDGRPC(c, ID)
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusOK, err.Error())
	}

	txts, err := art.GetText()
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusOK, err.Error())
	}

	imgs, vids := art.GetAssetsInfo()

	return c.Render(http.StatusOK, "article.htm", map[string]interface{}{
		"page":        "article_id",
		"assets_path": os.Getenv("DATA_ASSETS_PATH"),
		"id":          ID,
		"article":     art,
		"text":        txts,
		"imgs":        imgs,
		"vids":        vids,
	})
}
