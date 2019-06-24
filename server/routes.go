package server

import (
	"github.com/gorilla/mux"
	"github.com/reynaldipane/toko-ijah/order"
	"github.com/reynaldipane/toko-ijah/product"
	"github.com/reynaldipane/toko-ijah/purchase"
	"github.com/reynaldipane/toko-ijah/report"
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

	router.HandleFunc("/reports/product-values", report.GenerateProductValuesReport).Methods("GET")
	router.HandleFunc("/reports/sales-detail", report.GenerateProductSalesDetailReport).Methods("GET")

	router.HandleFunc("/reports/product-values/export!", report.ExportProductValuesReportCsv).Methods("GET")
	router.HandleFunc("/reports/sales-detail/export!", report.ExportSalesReportCsv).Methods("GET")

	router.HandleFunc("/report/downloads/{fileName}", report.DownloadProductValuesReport)
	return router
}
