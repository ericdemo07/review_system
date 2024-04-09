package mapper

import (
	"database/sql"
	"review_system/database/db_model"
	"review_system/model"
)

func CategoryToCategoryDBModelMapper(category model.Category) db_model.Category {
	var categoryRemapped db_model.Category

	categoryRemapped.Title = category.Title
	categoryRemapped.Text = sql.NullString{String: category.Text, Valid: true}
	categoryRemapped.Images = category.Images
	categoryRemapped.Tier = category.Tier

	return categoryRemapped
}

func ItemToItemDBModelMapper(item model.Item) db_model.Item {
	var itemRemapped db_model.Item

	itemRemapped.Title = item.Title
	itemRemapped.Text = item.Text
	itemRemapped.Images = item.Images
	itemRemapped.CategoryId = item.CategoryId

	return itemRemapped
}

func ReviewToReviewDBModelMapper(review model.Review) db_model.Review {
	var reviewRemapped db_model.Review

	reviewRemapped.Title = sql.NullString{String: review.Title, Valid: true}
	reviewRemapped.Text = review.Text
	reviewRemapped.Images = review.Images
	reviewRemapped.Rating = review.Rating
	reviewRemapped.ItemId = review.ItemId

	return reviewRemapped
}
