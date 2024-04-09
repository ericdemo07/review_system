package store

import (
	"errors"
	"gorm.io/gorm"
	"review_system/database/db_model"
	"review_system/model"
)

type CategoryStore struct {
	db *gorm.DB
}

func (cs *CategoryStore) IsCategoryExists(u uint) (*model.Category, error) {
	//TODO implement me
	panic("implement me")
}

func NewCategoryStore(db *gorm.DB) *CategoryStore {
	return &CategoryStore{
		db: db,
	}
}

func (cs *CategoryStore) CreateCategory(category *db_model.Category) error {
	err := cs.db.Create(&category)

	return err.Error
}

func (cs *CategoryStore) GetCategory(id uint) (*model.Category, error) {
	var m model.Category

	err := cs.db.Where(&model.Category{Id: id}).Find(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &m, err
}
