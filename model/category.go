package model

type Category struct {
	Id     uint     `json:"id"`
	Title  string   `json:"title"`
	Text   string   `json:"text"`
	Images []string `json:"images"`
	Tier   uint8    `json:"tier"`
}
