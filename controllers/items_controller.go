package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Kungfucoding23/bookstore_items_api/domain/items"
	"github.com/Kungfucoding23/bookstore_items_api/domain/queries"
	"github.com/Kungfucoding23/bookstore_items_api/services"
	"github.com/Kungfucoding23/bookstore_items_api/utils/http_utils"
	"github.com/Kungfucoding23/bookstore_utils-go/rest_errors"
	"github.com/gorilla/mux"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	// if err := oauth.AuthenticateRequest(r); err != nil {
	// 	restErr := rest_errors.NewBadRequestError("invalid request body")
	// 	http_utils.RespondError(w, restErr)
	// 	return
	// }
	// sellerID := oauth.GetCallerID(r)
	// if sellerID == 0 {
	// 	respErr := rest_errors.NewUnauthorizedError("invalid access token")
	// 	http_utils.RespondError(w, respErr)
	// 	return
	// }

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, respErr)
		return
	}

	// itemRequest.Seller = sellerID
	result, CreateErr := services.ItemsService.Create(itemRequest)
	if err != nil {
		http_utils.RespondError(w, CreateErr)
		return
	}
	http_utils.RespondJSON(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := strings.TrimSpace(vars["id"])
	item, err := services.ItemsService.Get(itemID)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}
	http_utils.RespondJSON(w, http.StatusOK, item)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, apiErr)
		return
	}
	defer r.Body.Close()
	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, apiErr)
		return
	}
	items, searchErr := services.ItemsService.Search(query)
	if searchErr != nil {
		http_utils.RespondError(w, searchErr)
		return
	}
	http_utils.RespondJSON(w, http.StatusOK, items)
}
