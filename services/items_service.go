package services

import (
	"github.com/JCFlores93/bookstore_items-api/domain/items"
	"github.com/JCFlores93/bookstore_utils_go/rest_errors"
	"net/http"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsService struct{}

type itemsServiceInterface interface {
	Create(item items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}

func (s *itemsService) Create(items.Item) (*items.Item, *rest_errors.RestErr) {
	return nil, &rest_errors.RestErr{
		Status:  http.StatusNotImplemented,
		Message: "implement me!",
		Error:   "not_implemented",
	}
}

func (s *itemsService) Get(string) (*items.Item, *rest_errors.RestErr) {
	return nil, &rest_errors.RestErr{
		Status:  http.StatusNotImplemented,
		Message: "implement me!",
		Error:   "not_implemented",
	}
}
