package controllers

import (
	"fmt"
	"net/http"

	"github.com/tony-landreth/bookstore_items-api/domain/items"
)

var (
	ItemsController itemsControllerInterface = &itemsController
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oath.AuthenticateRequest(r); err != nil {
		// TODO: Return error to user.
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		//TODO: return error json to user.
	}

	fmt.Println(result)
	//TODO: Return created item as JSON with HTTP status 201 - Created.
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hey")
}
