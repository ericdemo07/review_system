package item

import "review_system/model"

type Store interface {
	CreateItem(*model.Item) error
}
