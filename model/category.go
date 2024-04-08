package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Id     uint           `json:"id"`
	Title  string         `json:"title"`
	Text   string         `json:"text"`
	Images pq.StringArray `gorm:"type:string[]"`
	Tier   uint           `json:"tier"`
}
