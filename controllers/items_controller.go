package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/JCFlores93/bookstore_items-api/domain/items"
	"github.com/JCFlores93/bookstore_items-api/services"
	"github.com/JCFlores93/bookstore_items-api/utils/http_utils"
	"github.com/JCFlores93/bookstore_oauth_go/oauth"
	"github.com/JCFlores93/bookstore_utils_go/rest_errors"
	"io/ioutil"
	"net/http"
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
		//http_utils.RespondError(w, *err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, *respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, *respErr)
		return
	}
	itemRequest.Seller = oauth.GetClientId(r)

	//item := items.Item{
	//	Seller: oauth.GetClientId(r),
	//}
	result, createErr := services.ItemsService.Create(itemRequest)
	if err != nil {
		http_utils.RespondError(w, *createErr)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, result)
	fmt.Println(result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
