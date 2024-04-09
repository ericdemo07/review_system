package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"review_system/mapper"
	"review_system/model"
)

func (h *Handler) AddReview(c echo.Context) error {
	review := model.Review{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&review)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	dbReview := mapper.ReviewToReviewDBModelMapper(review)

	err = h.reviewStore.CreateReview(&dbReview)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	log.Printf("this is your review %#v", review)
	return c.String(http.StatusOK, "We got your Review!!!")
}
