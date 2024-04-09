package item

import (
	"review_system/database/db_model"
)

type Store interface {
	CreateItem(*db_model.Item) error
}
