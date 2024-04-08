package store

import (
	"gorm.io/gorm"
	"review_system/model"
)

type ItemStore struct {
	db *gorm.DB
}

func NewItemStore(db *gorm.DB) *ItemStore {
	return &ItemStore{
		db: db,
	}
}

func (cs *ItemStore) CreateItem(item *model.Item) error {
	err := cs.db.Create(&item)
	return err.Error
}
