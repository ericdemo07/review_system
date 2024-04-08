package main

import (
	"review_system/database"
	"review_system/handler"
	"review_system/router"
	"review_system/store"
)

func main() {
	r := router.New()
	g := r.Group("/api")

	db := database.New()
	database.AutoMigrate(db)

	cs := store.NewCategoryStore(db)
	is := store.NewItemStore(db)

	h := handler.NewHandler(cs, is)
	h.Register(g)

	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
}
