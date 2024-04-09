package model

type Review struct {
	Id     uint     `json:"id"`
	Title  string   `json:"title"`
	Text   string   `json:"text"`
	Rating uint     `json:"rating"`
	Images []string `json:"images"`
	ItemId uint8    `json:"item_id"`
}
