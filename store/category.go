package store

import (
	"gorm.io/gorm"
	"review_system/model"
)

type CategoryStore struct {
	db *gorm.DB
}

func NewCategoryStore(db *gorm.DB) *CategoryStore {
	return &CategoryStore{
		db: db,
	}
}

func (cs *CategoryStore) CreateCategory(category *model.Category) error {
	err := cs.db.Create(&category)

	return err.Error
}
