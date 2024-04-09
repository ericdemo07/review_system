package db_model

import (
	"database/sql"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

//https://gorm.io/docs/models.html

type Category struct {
	gorm.Model
	ID        uint           `gorm:"primaryKey"`
	Title     string         `gorm:"size:1000"`
	Text      sql.NullString `gorm:"size:65535"`
	Images    pq.StringArray `gorm:"type:string[]"`
	Tier      uint8          `gorm:"constraint:Max:3"`
	CreatedAt int64          `gorm:"autoUpdateTime:milli"`
	UpdatedAt int64          `gorm:"autoUpdateTime:milli"`
}
