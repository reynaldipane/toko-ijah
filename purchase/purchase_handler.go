package purchase

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	valid "github.com/asaskevich/govalidator"
	"github.com/reynaldipane/test-ijah/appcontext"
	helper "github.com/reynaldipane/test-ijah/helpers"
	productstock "github.com/reynaldipane/test-ijah/product_stock"
	purchasehistory "github.com/reynaldipane/test-ijah/purchase_history"
)

var purchaseService serviceInterface

func getService() serviceInterface {
	if purchaseService == nil {
		return &Service{
			repo: initRepository(appcontext.GetDB()),
			productStockService: &productstock.Service{
				Repo: productstock.InitRepository(appcontext.GetDB()),
			},
			purchaseHistoryService: &purchasehistory.Service{
				Repo: purchasehistory.InitRepository(appcontext.GetDB()),
			},
		}
	}

	return purchaseService
}

//CreatePurchase will handle creation of new purchase product
func CreatePurchase(w http.ResponseWriter, r *http.Request) {
	var purchase Purchase

	err := json.NewDecoder(r.Body).Decode(&purchase)

	if err != nil {
		helper.BuildResponseWithError(w, helper.ContentJSON, 400, err.Error())
		return
	}

	_, err = valid.ValidateStruct(purchase)

	if err != nil {
		helper.BuildResponseWithError(w, helper.ContentJSON, 400, err.Error())
		return
	}

	result, err := getService().createPurchase(purchase)

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

/*UpdatePurchase will handle uodate of existing purchase product
Update limited to the number received only
*/
func UpdatePurchase(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	purchaseID := vars["id"]

	var updatePurchase Purchase

	err := json.NewDecoder(r.Body).Decode(&updatePurchase)

	if err != nil {
		helper.BuildResponseWithError(w, helper.ContentJSON, 400, err.Error())
		return
	}

	_, err = valid.ValidateStruct(updatePurchase)

	if err != nil {
		helper.BuildResponseWithError(w, helper.ContentJSON, 400, err.Error())
		return
	}

	result, err := getService().updateNumberReceivedByID(purchaseID, updatePurchase)

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
