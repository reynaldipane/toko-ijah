package main

import (
	"log"
	"net/http"

	"github.com/reynaldipane/toko-ijah/product"
	productstock "github.com/reynaldipane/toko-ijah/product_stock"
	"github.com/reynaldipane/toko-ijah/server"

	"github.com/reynaldipane/toko-ijah/appcontext"
)

func main() {
	appcontext.InitContext()

	appcontext.GetDB().AutoMigrate(&product.Product{})
	appcontext.GetDB().AutoMigrate(&productstock.ProductStock{})

	router := server.CreateRouter()
	log.Fatal(http.ListenAndServe(":9000", router))
}
