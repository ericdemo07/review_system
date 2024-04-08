package category

import "review_system/model"

//this is kind of interface

type Store interface {
	CreateCategory(*model.Category) error
}
