package store

import (
	"gorm.io/gorm"
	"review_system/database/db_model"
)

type ItemStore struct {
	db *gorm.DB
}

func NewItemStore(db *gorm.DB) *ItemStore {
	return &ItemStore{
		db: db,
	}
}

func (cs *ItemStore) CreateItem(item *db_model.Item) error {
	err := cs.db.Create(&item)
	return err.Error
}
