package items

import (
	"errors"
	"fmt"

	"github.com/federicoleon/bookstore_utils-go/rest_errors"
	"github.com/tony-landreth/bookstore_items-api/clients/elasticsearch"
)

const (
	indexItems = "items"
	typeItem   = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {
	fmt.Println("try it")
	fmt.Println(elasticsearch.Client)
	result, err := elasticsearch.Client.Index(indexItems, typeItem, *i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}
