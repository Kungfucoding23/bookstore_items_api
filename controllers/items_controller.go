package controllers

import (
	"fmt"
	"net/http"

	"github.com/Kungfucoding23/bookstore_items_api/domain/items"
	"github.com/Kungfucoding23/bookstore_items_api/services"
	"github.com/Kungfucoding23/bookstore_oauth-go/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO: return error to the caller
	}
	item := items.Item{
		Seller: oauth.GetCallerID(r),
	}
	result, err := services.ItemsService.Create(item)
	if err != nil {
		//TODO: return error json to the user
	}
	fmt.Println(result)
	//TODO: return created item as json with HTTP status 201 created.
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
