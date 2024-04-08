package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"review_system/model"
)

func (h *Handler) AddItem(c echo.Context) error {
	item := model.Item{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&item)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	err = h.itemStore.CreateItem(&item)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	log.Printf("this is your item %#v", item)
	return c.String(http.StatusOK, "We got your Item!!!")
}
