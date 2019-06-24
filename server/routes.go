package server

import (
	"github.com/gorilla/mux"
	"github.com/reynaldipane/toko-ijah/order"
	"github.com/reynaldipane/toko-ijah/product"
	"github.com/reynaldipane/toko-ijah/purchase"
)

/*
CreateRouter will return a pointer to mux.Router
then will be used for as the router handler for app
*/
func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/products", product.CreateProductHandler).Methods("POST")

	router.HandleFunc("/purchases", purchase.CreatePurchase).Methods("POST")
	router.HandleFunc("/purchases/{id}", purchase.UpdatePurchase).Methods("PUT")

	router.HandleFunc("/orders", order.CreateOrder).Methods("POST")
	return router
}
