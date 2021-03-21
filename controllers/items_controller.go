package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Kungfucoding23/bookstore_items_api/domain/items"
	"github.com/Kungfucoding23/bookstore_items_api/services"
	"github.com/Kungfucoding23/bookstore_items_api/utils/http_utils"
	"github.com/Kungfucoding23/bookstore_oauth-go/oauth"
	"github.com/Kungfucoding23/bookstore_utils-go/rest_errors"
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
		restErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, restErr)
		return
	}

	var itemRequest items.Item

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = oauth.GetClientID(r)
	result, CreateErr := services.ItemsService.Create(itemRequest)
	if err != nil {
		http_utils.RespondError(w, CreateErr)
		return
	}
	http_utils.RespondJSON(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
