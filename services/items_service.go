package services

import (
	"net/http"

	"github.com/federicoleon/bookstore_utils-go/rest_errors"
	"github.com/tony-landreth/bookstore_items-api/domain/items"
)

var(
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(item items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(items.Item) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError(message: "implement me!", http.StatusNotImplemented, err: "not_implemented", causes: nil)
}

func (s *itemsService) Get(string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError(message: "implement me!", http.StatusNotImplemented, err: "not_implemented", causes: nil)
}
