package models

type Article struct {
	ID          int64  `json:"ID"`
	Timestamp   string `json:"Timestamp"`
	ID2         string `json:"ID2"`
	AID         string `json:"AID"`
	Title       string `json:"Title"`
	Category    string `json:"Category"`
	Abstract    string `json:"Abstract"`
	ArticleTags string `json:"ArticleTags"`
	Authors     string `json:"Authors"`
	Language    string `json:"Language"`
	Text        string `json:"Text"`
	Image       string `json:"Image"`
	Video       string `json:"Video"`
}
