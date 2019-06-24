package server

import (
	"github.com/gorilla/mux"
	"github.com/reynaldipane/toko-ijah/product"
)

/*
CreateRouter will return a pointer to mux.Router
then will be used for as the router handler for app
*/
func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/products", product.CreateProductHandler).Methods("POST")
	return router
}
