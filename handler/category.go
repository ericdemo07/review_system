package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"review_system/model"
)

func (h *Handler) AddCategory(c echo.Context) error {
	category := model.Category{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&category)

	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	err = h.categoryStore.CreateCategory(&category)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	log.Printf("this is yout category %#v", category)
	return c.String(http.StatusOK, "We got your Category!!!")
}
