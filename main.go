package main

import (
	"log"
	"net/http"

	"github.com/reynaldipane/toko-ijah/server"

	"github.com/reynaldipane/toko-ijah/appcontext"
)

func main() {
	appcontext.InitContext()

	router := server.CreateRouter()
	log.Fatal(http.ListenAndServe(":9000", router))
}
