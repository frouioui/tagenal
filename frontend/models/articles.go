package models

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Article struct {
	ID          int64  `json:"id"`
	Timestamp   int64  `json:"timestamp"`
	AID         string `json:"aid"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Abstract    string `json:"abstract"`
	ArticleTags string `json:"article_tags"`
	Authors     string `json:"authors"`
	Language    string `json:"language"`
	Text        string `json:"text"`
	Image       string `json:"image"`
	Video       string `json:"video"`
}

func (art *Article) GetAssetsInfo() ([]string, []string) {
	staticPath := os.Getenv("STATIC_PATH")
	imgs := strings.Split(art.Image, ",")
	vids := strings.Split(art.Video, ",")

	if art.Video == "" {
		vids = []string{}
	}
	if art.Image == "" {
		imgs = []string{}
	}
	for i := 0; i < len(imgs); i++ {
		imgs[i] = fmt.Sprintf("%s/articles/article%d/%s", staticPath, art.ID, imgs[i])
	}
	for i := 0; i < len(vids); i++ {
		vids[i] = fmt.Sprintf("%s/articles/article%d/%s", staticPath, art.ID, vids[i])
	}
	return imgs, vids
}

func (art *Article) GetText() (string, error) {
	fc, err := ioutil.ReadFile(fmt.Sprintf("%s/articles/article%d/%s", os.Getenv("DATA_ASSETS_PATH"), art.ID, art.Text))
	if err != nil {
		return "", err
	}
	return string(fc), nil
}
