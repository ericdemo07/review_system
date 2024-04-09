package db_model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID         uint           `gorm:"primaryKey"`
	Title      string         `gorm:"size:1000"`
	Text       string         `gorm:"size:65535"`
	Images     pq.StringArray `gorm:"type:string[]"`
	CreatedAt  int64          `gorm:"autoUpdateTime:milli"`
	UpdatedAt  int64          `gorm:"autoUpdateTime:milli"`
	CategoryId uint           `gorm:"TYPE:integer REFERENCES categories"`
}
