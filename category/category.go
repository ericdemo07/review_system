package category

import (
	"review_system/database/db_model"
	"review_system/model"
)

//this is kind of interface

type Store interface {
	CreateCategory(*db_model.Category) error
	GetCategory(uint) (*model.Category, error)
	IsCategoryExists(uint) (*model.Category, error)
}
