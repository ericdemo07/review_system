package review

import "review_system/database/db_model"

//this is kind of interface

type Store interface {
	CreateReview(*db_model.Review) error
}
