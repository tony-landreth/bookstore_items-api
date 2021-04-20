package app

import "github.com/tony-landreth/bookstore_items-api/controllers"

func mapUrls() {
	router.HandleFunc(path: "/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	controllers.UsersController.Create()
}
