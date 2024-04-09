package model

type Item struct {
	Id         uint     `json:"id"`
	Title      string   `json:"title"`
	Text       string   `json:"text"`
	Images     []string `json:"images"`
	CategoryId uint     `json:"category_id"`
}
