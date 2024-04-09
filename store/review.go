package store

import (
	"gorm.io/gorm"
	"review_system/database/db_model"
)

type ReviewStore struct {
	db *gorm.DB
}

func NewReviewStore(db *gorm.DB) *ReviewStore {
	return &ReviewStore{
		db: db,
	}
}

func (cs *ReviewStore) CreateReview(review *db_model.Review) error {
	err := cs.db.Create(&review)
	return err.Error
}
