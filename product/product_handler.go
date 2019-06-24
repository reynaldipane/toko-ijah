package product

import (
	"encoding/json"
	"net/http"

	productstock "github.com/reynaldipane/test-ijah/product_stock"

	"github.com/reynaldipane/test-ijah/appcontext"

	helper "github.com/reynaldipane/test-ijah/helpers"

	valid "github.com/asaskevich/govalidator"
)

var productService serviceInterface

func getService() serviceInterface {
	if productService == nil {
		productService = &Service{
			repo: initRepository(appcontext.GetDB()),
			productStockService: &productstock.Service{
				Repo: productstock.InitRepository(appcontext.GetDB()),
			},
		}
	}

	return productService
}

//CreateProductHandler will handle the creation of new product
func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		helper.BuildResponseWithError(w, helper.ContentJSON, 400, err.Error())
		return
	}

	_, err = valid.ValidateStruct(product)

	if err != nil {
		helper.BuildResponseWithError(w, helper.ContentJSON, 400, err.Error())
		return
	}

	result, err := getService().createProduct(product)

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
