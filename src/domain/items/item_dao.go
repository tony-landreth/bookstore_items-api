package items

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/federicoleon/bookstore_utils-go/rest_errors"
	"github.com/tony-landreth/bookstore_items-api/src/clients/elasticsearch"
	"github.com/tony-landreth/bookstore_items-api/src/domain/queries"
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

func (i *Item) Search(query queries.EsQuery) ([]Item, rest_errors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}

	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying to parse response", errors.New("database error"))
		}
		item.Id = hit.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundError("no items found matching given criteria")
	}

	return items, nil
}
