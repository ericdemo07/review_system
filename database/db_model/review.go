package db_model

import (
	"database/sql"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ID        uint           `gorm:"primaryKey"`
	Title     sql.NullString `gorm:"size:1000"`
	Text      string         `gorm:"size:65535"`
	Rating    uint           `gorm:"constraint:Min:1 constraint:Max:5"`
	Images    pq.StringArray `gorm:"type:string[]"`
	CreatedAt int64          `gorm:"autoUpdateTime:milli"`
	UpdatedAt int64          `gorm:"autoUpdateTime:milli"`
	ItemId    uint8          `gorm:"TYPE:integer REFERENCES items"`
}
