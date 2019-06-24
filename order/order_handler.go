package order

import (
	"encoding/json"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/reynaldipane/toko-ijah/appcontext"
	helper "github.com/reynaldipane/toko-ijah/helpers"
	productstock "github.com/reynaldipane/toko-ijah/product_stock"
)

var orderService ServiceInterface

func getService() ServiceInterface {
	if orderService == nil {
		return &Service{
			Repo: InitRepository(appcontext.GetDB()),
			productStockService: &productstock.Service{
				Repo: productstock.InitRepository(appcontext.GetDB()),
			},
		}
	}

	return orderService
}

//CreateOrder will handler order creation
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order Order

	err := json.NewDecoder(r.Body).Decode(&order)

	if err != nil {
		helper.BuildResponseWithError(w, helper.ContentJSON, 400, err.Error())
		return
	}

	_, err = valid.ValidateStruct(order)

	if err != nil {
		helper.BuildResponseWithError(w, helper.ContentJSON, 400, err.Error())
		return
	}

	result, err := getService().createOrder(order)

	if err != nil {
		helper.BuildResponseWithError(w, helper.ContentJSON, 500, err.Error())
		return
	}

	response, err := json.Marshal(result)

	if err != nil {
		helper.BuildResponseWithError(w, helper.ContentJSON, 500, err.Error())
		return
	}

	helper.BuildResponse(w, helper.ContentJSON, 201, string(response))
}
