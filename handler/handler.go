package handler

import (
	"review_system/category"
	"review_system/item"
	"review_system/review"
)

type Handler struct {
	categoryStore category.Store
	itemStore     item.Store
	reviewStore   review.Store
}

func NewHandler(cs category.Store, is item.Store, rs review.Store) *Handler {
	return &Handler{
		categoryStore: cs,
		itemStore:     is,
		reviewStore:   rs,
	}
}
