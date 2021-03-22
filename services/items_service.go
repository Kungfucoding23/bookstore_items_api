package services

import (
	"github.com/Kungfucoding23/bookstore_items_api/domain/items"
	"github.com/Kungfucoding23/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, rest_errors.NewRestError(err.Message(), err.Status(), err.Error(), err.Causes())
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{ID: id}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}
