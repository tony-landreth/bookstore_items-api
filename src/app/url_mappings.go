package app

import (
	"net/http"

	"github.com/tony-landreth/bookstore_items-api/src/controllers"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/search", controllers.ItemsController.Search).Methods(http.MethodPost)
}
