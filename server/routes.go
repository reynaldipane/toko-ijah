package server

import (
	"github.com/gorilla/mux"
)

/*
CreateRouter will return a pointer to mux.Router
then will be used for as the router handler for app
*/
func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}
