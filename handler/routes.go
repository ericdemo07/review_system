package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) Register(g *echo.Group) {
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	category := g.Group("/category")
	category.POST("", h.AddCategory)

	item := g.Group("/item")
	item.POST("", h.AddItem)
}
