package handler

import (
	"review_system/category"
	"review_system/item"
)

type Handler struct {
	categoryStore category.Store
	itemStore     item.Store
}

func NewHandler(cs category.Store, is item.Store) *Handler {
	return &Handler{
		categoryStore: cs,
		itemStore:     is,
	}
}
